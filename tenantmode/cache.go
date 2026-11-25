package tenantmode

import (
	"context"
	"log"
	"maps"
	"sync/atomic"
	"time"
)

type snapshot struct {
	entries map[string]Entry
	at      time.Time
}

// Cache is the in-memory registry every replica reads for one module. It
// never starts empty and keeps its last good snapshot when a reload fails
// (fail closed).
type Cache struct {
	module Module
	load   func(context.Context) (map[string]Entry, error)
	snap   atomic.Pointer[snapshot]
}

// NewCache loads the module's registry once. On error it returns no cache;
// the caller retries before serving anything.
func NewCache(ctx context.Context, s *Store) (*Cache, error) {
	return newCache(ctx, s.module, s.LoadAll)
}

func newCache(ctx context.Context, module Module, load func(context.Context) (map[string]Entry, error)) (*Cache, error) {
	m, err := load(ctx)
	if err != nil {
		return nil, err
	}
	c := &Cache{module: module, load: load}
	c.snap.Store(&snapshot{entries: m, at: time.Now()})
	return c, nil
}

// Get returns the tenant's entry, or legacy for an unknown domain. It never
// touches the database.
func (c *Cache) Get(domain string) Entry {
	d := normalise(domain)
	e, ok := c.snap.Load().entries[d]
	if !ok {
		return Entry{Domain: d, Module: c.module, Mode: ModeLegacy}
	}
	e.JobsEnabled = maps.Clone(e.JobsEnabled)
	return e
}

// Snapshot returns a copy of every registered entry (tenants with no
// document, i.e. plain legacy, are absent). It never touches the database.
func (c *Cache) Snapshot() map[string]Entry {
	entries := c.snap.Load().entries
	m := make(map[string]Entry, len(entries))
	for d, e := range entries {
		e.JobsEnabled = maps.Clone(e.JobsEnabled)
		m[d] = e
	}
	return m
}

// LoadedAt is when the current snapshot was loaded.
func (c *Cache) LoadedAt() time.Time { return c.snap.Load().at }

// Run reloads every ReloadInterval until ctx ends.
func (c *Cache) Run(ctx context.Context) {
	t := time.NewTicker(ReloadInterval)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			c.reload(ctx)
		}
	}
}

func (c *Cache) reload(ctx context.Context) {
	m, err := c.load(ctx)
	if err != nil {
		log.Printf("[tenantmode %s] reload failed, keeping snapshot from %s: %v", c.module, c.snap.Load().at.UTC().Format(time.RFC3339), err)
		return
	}
	c.snap.Store(&snapshot{entries: m, at: time.Now()})
}
