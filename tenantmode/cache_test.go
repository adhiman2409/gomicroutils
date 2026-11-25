package tenantmode

import (
	"context"
	"errors"
	"testing"
)

func TestCache(t *testing.T) {
	ctx := context.Background()
	if _, err := newCache(ctx, ModuleTaskMgr, func(context.Context) (map[string]Entry, error) { return nil, errors.New("down") }); err == nil {
		t.Fatal("NewCache over a failing loader returned a cache")
	}

	next := map[string]Entry{"acme.test": {Domain: "acme.test", Module: ModuleTaskMgr, Mode: ModeV1, JobsEnabled: map[string]bool{"outbox": true}}}
	var fail bool
	c, err := newCache(ctx, ModuleTaskMgr, func(context.Context) (map[string]Entry, error) {
		if fail {
			return nil, errors.New("down")
		}
		return next, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if c.Get("ACME.test").Mode != ModeV1 {
		t.Error("Get is not case-insensitive")
	}
	if e := c.Get("unknown.test"); e.Mode != ModeLegacy || e.Domain != "unknown.test" || e.Module != ModuleTaskMgr {
		t.Errorf("unknown domain: %+v", e)
	}
	c.Get("acme.test").JobsEnabled["outbox"] = false
	if !c.Get("acme.test").JobsEnabled["outbox"] {
		t.Error("a caller mutated the shared snapshot")
	}

	fail = true
	c.reload(ctx)
	if c.Get("acme.test").Mode != ModeV1 {
		t.Error("a failed reload dropped the snapshot")
	}

	fail = false
	next = map[string]Entry{"acme.test": {Domain: "acme.test", Mode: ModeLegacy}}
	c.reload(ctx)
	if c.Get("acme.test").Mode != ModeLegacy {
		t.Error("a good reload did not replace the snapshot")
	}
}
