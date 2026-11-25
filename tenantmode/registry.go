// Package tenantmode is the shared tenant-mode registry: which world (legacy,
// migrating or v1) each tenant lives in, per rebuilt module. It is shared
// because several services must see the same answer: for taskmgr, kyc freezes
// legacy writes and attendance, employee and helpdesk route calls by it.
//
// The registry is one collection in the master database, one document per
// (domain, module). Services read it through a Cache that reloads every
// ReloadInterval and never touches the database on a request. Only an
// operator CLI writes it, through Store.Apply.
//
// The rules are the ones FAMS uses in inventory, unchanged: legacy ->
// migrating -> v1 -> final, rollback to legacy from migrating or from a v1
// tenant that is not final, and a test-tenant flag that enables v1 routes on
// a legacy tenant listed in the module's test-tenant allow-list.
package tenantmode

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/adhiman2409/gomicroutils/hrms"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Module names a rebuilt module that has its own mode per tenant.
type Module string

const ModuleTaskMgr Module = "taskmgr"

type Mode string

const (
	ModeLegacy    Mode = "legacy"
	ModeMigrating Mode = "migrating"
	ModeV1        Mode = "v1"
)

// ReloadInterval is how often every replica reloads the registry.
const ReloadInterval = 15 * time.Second

var (
	ErrTransitionNotAllowed = errors.New("transition not allowed")
	ErrNotTestTenant        = errors.New("domain is not in this module's test-tenant list")
	ErrConcurrentChange     = errors.New("the tenant's mode was changed by someone else; read it again")
	ErrInvalidChange        = errors.New("a change must set exactly one of: mode, rollback, final, test tenant, job")
)

// Entry is one tenant's state for one module. A tenant with no document is
// legacy.
type Entry struct {
	Domain         string
	Module         Module
	Mode           Mode
	Final          bool
	TestTenant     bool
	JobsEnabled    map[string]bool
	MigrationRunID string
	ChangedAt      time.Time
	ChangedBy      string
}

// Enabled reports whether v1 routes answer for the tenant.
func (e Entry) Enabled() bool {
	return e.Mode == ModeV1 || (e.TestTenant && e.Mode == ModeLegacy)
}

// LegacyWritable reports whether legacy writes are allowed. It looks at the
// mode alone: a test tenant that is migrating freezes like a real one.
func (e Entry) LegacyWritable() bool { return e.Mode == ModeLegacy }

// Reader answers from memory. An unknown domain is legacy.
type Reader interface {
	Get(domain string) Entry
}

// Change is one requested registry change. Exactly one field group is set.
type Change struct {
	SetMode        *Mode // migrating or v1
	Rollback       bool  // migrating -> legacy, or v1 (not final) -> legacy
	Final          bool  // v1 -> final
	TestTenant     *bool
	Job            string // with JobOn
	JobOn          bool
	MigrationRunID string
}

func normalise(domain string) string { return strings.ToLower(strings.TrimSpace(domain)) }

// TestTenantsFromEnv reads a comma-separated domain list from the named
// environment variable, e.g. TASKMGR_TEST_TENANTS.
func TestTenantsFromEnv(name string) map[string]bool {
	out := map[string]bool{}
	for _, d := range strings.Split(os.Getenv(name), ",") {
		if d = normalise(d); d != "" {
			out[d] = true
		}
	}
	return out
}

// plan checks a change against the transition table and returns the next
// entry and a history label. It touches nothing.
func plan(cur Entry, c Change, testTenantAllowed bool, jobs []string) (Entry, string, error) {
	set := 0
	for _, b := range []bool{c.SetMode != nil, c.Rollback, c.Final, c.TestTenant != nil, c.Job != ""} {
		if b {
			set++
		}
	}
	if set != 1 {
		return cur, "", ErrInvalidChange
	}
	next := cur
	next.JobsEnabled = maps.Clone(cur.JobsEnabled)
	if next.JobsEnabled == nil {
		next.JobsEnabled = map[string]bool{}
	}
	refuse := func(what string) (Entry, string, error) {
		return cur, "", fmt.Errorf("%w: %s from %s%s", ErrTransitionNotAllowed, what, cur.Mode, map[bool]string{true: " (final)"}[cur.Final])
	}

	switch {
	case c.Job != "":
		if !slices.Contains(jobs, c.Job) {
			if len(jobs) == 0 {
				return cur, "", fmt.Errorf("unknown job %q; this module has no jobs", c.Job)
			}
			return cur, "", fmt.Errorf("unknown job %q; use one of %s", c.Job, strings.Join(jobs, ", "))
		}
		next.JobsEnabled[c.Job] = c.JobOn
		return next, fmt.Sprintf("job %s %s", c.Job, onOff(c.JobOn)), nil
	case cur.Final:
		return refuse("any change except jobs")
	case c.SetMode != nil:
		to := *c.SetMode
		if (cur.Mode == ModeLegacy && to == ModeMigrating) || (cur.Mode == ModeMigrating && to == ModeV1) {
			next.Mode = to
			if c.MigrationRunID != "" {
				next.MigrationRunID = c.MigrationRunID
			}
			return next, "set " + string(to), nil
		}
		return refuse("set " + string(to))
	case c.Rollback:
		if cur.Mode == ModeMigrating || cur.Mode == ModeV1 {
			next.Mode = ModeLegacy
			return next, "rollback", nil
		}
		return refuse("rollback")
	case c.Final:
		if cur.Mode == ModeV1 {
			next.Final = true
			return next, "final", nil
		}
		return refuse("final")
	default: // TestTenant
		if !testTenantAllowed {
			return cur, "", ErrNotTestTenant
		}
		next.TestTenant = *c.TestTenant
		return next, "test-tenant " + onOff(*c.TestTenant), nil
	}
}

func onOff(b bool) string {
	if b {
		return "on"
	}
	return "off"
}

// Option configures a Store.
type Option func(*Store)

// WithTestTenants sets the domains allowed to turn the test-tenant flag on.
func WithTestTenants(domains map[string]bool) Option {
	return func(s *Store) {
		s.testTenants = map[string]bool{}
		for d, ok := range domains {
			if ok {
				s.testTenants[normalise(d)] = true
			}
		}
	}
}

// WithJobs sets the job keys allowed in JobsEnabled for the module.
func WithJobs(jobs ...string) Option {
	return func(s *Store) { s.jobs = slices.Clone(jobs) }
}

// Store reads and writes one module's rows of the registry.
type Store struct {
	module      Module
	coll        *mongo.Collection
	testTenants map[string]bool
	jobs        []string
}

func masterDB(client *mongo.Client) *mongo.Database {
	return client.Database(strings.ReplaceAll(hrms.GetMasterDomain(), ".", "_"))
}

// NewStore returns the registry store for one module.
func NewStore(client *mongo.Client, module Module, opts ...Option) *Store {
	s := &Store{
		module:      module,
		coll:        masterDB(client).Collection(hrms.MASTER_COLLECTION_TENANT_MODES),
		testTenants: map[string]bool{},
	}
	for _, o := range opts {
		o(s)
	}
	return s
}

// Module returns the module this store reads and writes.
func (s *Store) Module() Module { return s.module }

// Jobs returns the job keys this module allows.
func (s *Store) Jobs() []string { return slices.Clone(s.jobs) }

// EnsureIndexes creates the unique (domain, module) index.
func (s *Store) EnsureIndexes(ctx context.Context) error {
	_, err := s.coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "domain", Value: 1}, {Key: "module", Value: 1}},
		Options: options.Index().SetUnique(true).SetName("tenant_modes_domain_module_unique"),
	})
	return err
}

func (s *Store) key(domain string) bson.D {
	return bson.D{{Key: "domain", Value: domain}, {Key: "module", Value: string(s.module)}}
}

// LoadAll reads every document of the module, keyed by lower-cased domain.
func (s *Store) LoadAll(ctx context.Context) (map[string]Entry, error) {
	cur, err := s.coll.Find(ctx, bson.D{{Key: "module", Value: string(s.module)}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	out := map[string]Entry{}
	for cur.Next(ctx) {
		e := entryFromRaw(cur.Current)
		e.Module = s.module
		out[e.Domain] = e
	}
	return out, cur.Err()
}

// Document returns the stored document as is (with its history), or nil.
func (s *Store) Document(ctx context.Context, domain string) (bson.M, error) {
	domain = normalise(domain)
	var doc bson.M
	err := s.coll.FindOne(ctx, s.key(domain)).Decode(&doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	return doc, err
}

// Apply checks the change and writes it with a compare-and-set on
// changed_at, so two operators can never both win.
func (s *Store) Apply(ctx context.Context, domain string, c Change, by string) (Entry, error) {
	domain = normalise(domain)
	if domain == "" || strings.TrimSpace(by) == "" {
		return Entry{}, errors.New("domain and changed_by are required")
	}
	cur := Entry{Domain: domain, Module: s.module, Mode: ModeLegacy}
	var raw bson.Raw
	err := s.coll.FindOne(ctx, s.key(domain)).Decode(&raw)
	exists := err == nil
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return Entry{}, err
	}
	if exists {
		cur = entryFromRaw(raw)
		cur.Module = s.module
	}

	next, label, err := plan(cur, c, s.testTenants[domain], s.jobs)
	if err != nil {
		return Entry{}, err
	}
	now := time.Now().UTC().Truncate(time.Millisecond)
	next.ChangedAt, next.ChangedBy = now, by

	state := func(e Entry) bson.M {
		return bson.M{"mode": e.Mode, "final": e.Final, "test_tenant": e.TestTenant, "jobs_enabled": e.JobsEnabled}
	}
	history := bson.M{"at": now, "by": by, "from": state(cur), "to": state(next), "change": label}
	fields := bson.M{
		"domain": domain, "module": string(s.module), "mode": next.Mode, "final": next.Final,
		"test_tenant": next.TestTenant, "jobs_enabled": next.JobsEnabled,
		"migration_run_id": next.MigrationRunID, "changed_at": now, "changed_by": by,
	}

	if !exists {
		fields["history"] = bson.A{history}
		if _, err := s.coll.InsertOne(ctx, fields); err != nil {
			if mongo.IsDuplicateKeyError(err) {
				return Entry{}, ErrConcurrentChange
			}
			return Entry{}, err
		}
		return next, nil
	}
	// Compare-and-set on changed_at. A document that never had one (written
	// by hand) matches on its absence instead.
	var stamp any = cur.ChangedAt
	if cur.ChangedAt.IsZero() {
		stamp = bson.M{"$exists": false}
	}
	filter := append(s.key(domain), bson.E{Key: "changed_at", Value: stamp})
	res, err := s.coll.UpdateOne(ctx, filter, bson.M{"$set": fields, "$push": bson.M{"history": history}})
	if err != nil {
		return Entry{}, err
	}
	if res.MatchedCount == 0 {
		return Entry{}, ErrConcurrentChange
	}
	return next, nil
}

// entryFromRaw maps a stored document to an Entry. Missing fields keep their
// zero value; a missing mode is legacy.
func entryFromRaw(r bson.Raw) Entry {
	e := Entry{JobsEnabled: map[string]bool{}}
	d, _ := r.Lookup("domain").StringValueOK()
	e.Domain = normalise(d)
	mod, _ := r.Lookup("module").StringValueOK()
	e.Module = Module(mod)
	m, _ := r.Lookup("mode").StringValueOK()
	e.Mode = Mode(m)
	if e.Mode == "" {
		e.Mode = ModeLegacy
	}
	e.Final, _ = r.Lookup("final").BooleanOK()
	e.TestTenant, _ = r.Lookup("test_tenant").BooleanOK()
	if jobs, ok := r.Lookup("jobs_enabled").DocumentOK(); ok {
		elems, _ := jobs.Elements()
		for _, el := range elems {
			on, _ := el.Value().BooleanOK()
			e.JobsEnabled[el.Key()] = on
		}
	}
	e.MigrationRunID, _ = r.Lookup("migration_run_id").StringValueOK()
	if t, ok := r.Lookup("changed_at").TimeOK(); ok {
		e.ChangedAt = t.UTC()
	}
	e.ChangedBy, _ = r.Lookup("changed_by").StringValueOK()
	return e
}
