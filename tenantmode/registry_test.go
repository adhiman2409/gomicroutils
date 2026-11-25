package tenantmode

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func modePtr(m Mode) *Mode { return &m }
func boolPtr(b bool) *bool { return &b }

var testJobs = []string{"outbox", "reminders"}

func TestPlan(t *testing.T) {
	legacy := Entry{Mode: ModeLegacy}
	migrating := Entry{Mode: ModeMigrating}
	v1 := Entry{Mode: ModeV1}
	final := Entry{Mode: ModeV1, Final: true}
	cases := []struct {
		name string
		cur  Entry
		c    Change
		want Entry
		err  error
	}{
		{"legacy to migrating", legacy, Change{SetMode: modePtr(ModeMigrating)}, Entry{Mode: ModeMigrating}, nil},
		{"migrating to v1", migrating, Change{SetMode: modePtr(ModeV1)}, Entry{Mode: ModeV1}, nil},
		{"v1 to final", v1, Change{Final: true}, Entry{Mode: ModeV1, Final: true}, nil},
		{"migrating rollback", migrating, Change{Rollback: true}, Entry{Mode: ModeLegacy}, nil},
		{"v1 rollback", v1, Change{Rollback: true}, Entry{Mode: ModeLegacy}, nil},
		{"legacy to v1", legacy, Change{SetMode: modePtr(ModeV1)}, Entry{}, ErrTransitionNotAllowed},
		{"v1 to migrating", v1, Change{SetMode: modePtr(ModeMigrating)}, Entry{}, ErrTransitionNotAllowed},
		{"set legacy", migrating, Change{SetMode: modePtr(ModeLegacy)}, Entry{}, ErrTransitionNotAllowed},
		{"legacy rollback", legacy, Change{Rollback: true}, Entry{}, ErrTransitionNotAllowed},
		{"final rollback", final, Change{Rollback: true}, Entry{}, ErrTransitionNotAllowed},
		{"final twice", final, Change{Final: true}, Entry{}, ErrTransitionNotAllowed},
		{"final test tenant", final, Change{TestTenant: boolPtr(true)}, Entry{}, ErrTransitionNotAllowed},
		{"migrating final", migrating, Change{Final: true}, Entry{}, ErrTransitionNotAllowed},
		{"two changes", legacy, Change{Final: true, Rollback: true}, Entry{}, ErrInvalidChange},
		{"no change", legacy, Change{}, Entry{}, ErrInvalidChange},
	}
	for _, c := range cases {
		got, _, err := plan(c.cur, c.c, true, testJobs)
		if !errors.Is(err, c.err) {
			t.Errorf("%s: err %v, want %v", c.name, err, c.err)
			continue
		}
		if err == nil && (got.Mode != c.want.Mode || got.Final != c.want.Final) {
			t.Errorf("%s: got %s final=%v", c.name, got.Mode, got.Final)
		}
	}

	if _, _, err := plan(legacy, Change{TestTenant: boolPtr(true)}, false, testJobs); !errors.Is(err, ErrNotTestTenant) {
		t.Errorf("test tenant outside allow-list: %v", err)
	}
	if _, _, err := plan(legacy, Change{Job: "payroll", JobOn: true}, true, testJobs); err == nil {
		t.Error("unknown job accepted")
	}
	if _, _, err := plan(legacy, Change{Job: "outbox", JobOn: true}, true, nil); err == nil || !strings.Contains(err.Error(), "no jobs") {
		t.Errorf("job on a module with no jobs: %v", err)
	}
	got, _, err := plan(final, Change{Job: "outbox", JobOn: true}, true, testJobs)
	if err != nil || !got.JobsEnabled["outbox"] {
		t.Errorf("job on a final tenant: %v %v", got.JobsEnabled, err)
	}
	if final.JobsEnabled != nil {
		t.Error("plan mutated the current entry")
	}
}

func TestEnabledAndLegacyWritable(t *testing.T) {
	for _, c := range []struct {
		mode           Mode
		test           bool
		enabled, write bool
	}{
		{ModeLegacy, false, false, true},
		{ModeLegacy, true, true, true},
		{ModeMigrating, false, false, false},
		{ModeMigrating, true, false, false},
		{ModeV1, false, true, false},
		{ModeV1, true, true, false},
	} {
		e := Entry{Mode: c.mode, TestTenant: c.test}
		if e.Enabled() != c.enabled || e.LegacyWritable() != c.write {
			t.Errorf("%s test=%v: enabled=%v writable=%v", c.mode, c.test, e.Enabled(), e.LegacyWritable())
		}
	}
}

func TestTestTenantsFromEnv(t *testing.T) {
	t.Setenv("X_TEST_TENANTS", " Demo.Example.com, ,other.test")
	got := TestTenantsFromEnv("X_TEST_TENANTS")
	if len(got) != 2 || !got["demo.example.com"] || !got["other.test"] {
		t.Errorf("got %v", got)
	}
}

// testClient connects to TENANTMODE_TEST_MONGO_URI and skips when it is unset.
// The store tests need a real server; the rest of this file does not.
func testClient(t *testing.T) *mongo.Client {
	t.Helper()
	uri := os.Getenv("TENANTMODE_TEST_MONGO_URI")
	if uri == "" {
		t.Skip("TENANTMODE_TEST_MONGO_URI not set")
	}
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = c.Disconnect(context.Background()) })
	return c
}

// testMaster points the master domain at a throwaway database and drops it
// at the end.
func testMaster(t *testing.T, c *mongo.Client) {
	t.Helper()
	master := fmt.Sprintf("ct-master-%d.test", os.Getpid())
	t.Setenv("MASTER_DOMAIN", master)
	t.Cleanup(func() { _ = masterDB(c).Drop(context.Background()) })
}

func TestStoreApply(t *testing.T) {
	c := testClient(t)
	testMaster(t, c)
	ctx := context.Background()
	s := NewStore(c, ModuleTaskMgr, WithJobs(testJobs...), WithTestTenants(map[string]bool{"Demo.Example.com": true}))
	if err := s.EnsureIndexes(ctx); err != nil {
		t.Fatal(err)
	}
	const d = "acme.test"

	for i, ch := range []Change{{SetMode: modePtr(ModeMigrating)}, {SetMode: modePtr(ModeV1)}, {Final: true}} {
		if _, err := s.Apply(ctx, "ACME.test", ch, "ops@example.com"); err != nil {
			t.Fatalf("step %d: %v", i, err)
		}
	}
	if _, err := s.Apply(ctx, d, Change{Rollback: true}, "ops@example.com"); !errors.Is(err, ErrTransitionNotAllowed) {
		t.Fatalf("rollback after final: %v", err)
	}
	doc, err := s.Document(ctx, d)
	if err != nil {
		t.Fatal(err)
	}
	if h := doc["history"].(bson.A); len(h) != 3 {
		t.Errorf("history has %d entries, want 3", len(h))
	}
	if doc["module"] != "taskmgr" {
		t.Errorf("module = %v", doc["module"])
	}
	all, err := s.LoadAll(ctx)
	if err != nil || all[d].Mode != ModeV1 || !all[d].Final || all[d].Module != ModuleTaskMgr {
		t.Errorf("LoadAll: %+v %v", all[d], err)
	}

	if _, err := s.Apply(ctx, "nottest.test", Change{TestTenant: boolPtr(true)}, "ops@example.com"); !errors.Is(err, ErrNotTestTenant) {
		t.Errorf("test tenant outside allow-list: %v", err)
	}
	if e, err := s.Apply(ctx, "demo.example.com", Change{TestTenant: boolPtr(true)}, "ops@example.com"); err != nil || !e.Enabled() {
		t.Errorf("test tenant: %+v %v", e, err)
	}

	// Another module's rows are invisible to this one, for the same domain.
	other := NewStore(c, Module("other"))
	if all, err := other.LoadAll(ctx); err != nil || len(all) != 0 {
		t.Errorf("other module sees taskmgr rows: %v %v", all, err)
	}
	if e, err := other.Apply(ctx, d, Change{SetMode: modePtr(ModeMigrating)}, "ops@example.com"); err != nil || e.Mode != ModeMigrating {
		t.Errorf("other module apply: %+v %v", e, err)
	}
	if all, _ := s.LoadAll(ctx); all[d].Mode != ModeV1 {
		t.Errorf("other module change leaked into taskmgr: %+v", all[d])
	}
}

func TestStoreConcurrentApply(t *testing.T) {
	c := testClient(t)
	testMaster(t, c)
	ctx := context.Background()
	s := NewStore(c, ModuleTaskMgr)
	if err := s.EnsureIndexes(ctx); err != nil {
		t.Fatal(err)
	}
	for _, d := range []string{"new.test", "old.test"} {
		next := ModeMigrating
		if d == "old.test" { // existing document: compare-and-set on update
			if _, err := s.Apply(ctx, d, Change{SetMode: modePtr(ModeMigrating)}, "a"); err != nil {
				t.Fatal(err)
			}
			next = ModeV1
		}
		var wg sync.WaitGroup
		errs := make([]error, 8)
		for i := range errs {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				_, errs[i] = s.Apply(ctx, d, Change{SetMode: &next}, "op")
			}(i)
		}
		wg.Wait()
		ok := 0
		for _, err := range errs {
			switch {
			case err == nil:
				ok++
			case errors.Is(err, ErrConcurrentChange), errors.Is(err, ErrTransitionNotAllowed):
			default:
				t.Errorf("%s: unexpected error %v", d, err)
			}
		}
		if ok != 1 {
			t.Errorf("%s: %d writers won, want 1", d, ok)
		}
	}
}
