package grpcauth

import (
	"context"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	keyA = strings.Repeat("a", MinKeyLength)
	keyK = strings.Repeat("k", MinKeyLength)
)

func testPolicy(t *testing.T) *Policy {
	t.Helper()
	p, err := NewPolicy(
		map[string]string{"attendance": keyA, "kyc": keyK},
		map[string][]string{
			"attendance": {"/compat.Compat/ListPendingTasks", "/compat.Compat/AddApprovedHours"},
			"kyc":        {"/compat.Compat/*"},
		})
	if err != nil {
		t.Fatal(err)
	}
	return p
}

func incoming(kv ...string) context.Context {
	return metadata.NewIncomingContext(context.Background(), metadata.Pairs(kv...))
}

func TestCheck(t *testing.T) {
	p := testPolicy(t)
	cases := []struct {
		name   string
		ctx    context.Context
		method string
		code   codes.Code
		caller string
	}{
		{"allowed method", incoming(CallerKeyHeader, keyA, DomainHeader, "Demo.Unirms.com"), "/compat.Compat/AddApprovedHours", codes.OK, "attendance"},
		{"whole service", incoming(CallerKeyHeader, keyK, DomainHeader, "demo.unirms.com"), "/compat.Compat/GetProjectHours", codes.OK, "kyc"},
		{"method not allowed", incoming(CallerKeyHeader, keyA, DomainHeader, "demo.unirms.com"), "/compat.Compat/GetProjectHours", codes.PermissionDenied, ""},
		{"other service", incoming(CallerKeyHeader, keyK, DomainHeader, "demo.unirms.com"), "/other.Svc/Get", codes.PermissionDenied, ""},
		{"wrong key", incoming(CallerKeyHeader, strings.Repeat("z", MinKeyLength), DomainHeader, "demo.unirms.com"), "/compat.Compat/AddApprovedHours", codes.Unauthenticated, ""},
		{"no key", incoming(DomainHeader, "demo.unirms.com"), "/compat.Compat/AddApprovedHours", codes.Unauthenticated, ""},
		{"no metadata", context.Background(), "/compat.Compat/AddApprovedHours", codes.Unauthenticated, ""},
		{"no domain", incoming(CallerKeyHeader, keyA), "/compat.Compat/AddApprovedHours", codes.InvalidArgument, ""},
	}
	for _, c := range cases {
		ctx, err := p.Check(c.ctx, c.method)
		if got := status.Code(err); got != c.code {
			t.Errorf("%s: code %s, want %s (%v)", c.name, got, c.code, err)
			continue
		}
		if err == nil {
			if Caller(ctx) != c.caller || Domain(ctx) != "demo.unirms.com" {
				t.Errorf("%s: caller %q domain %q", c.name, Caller(ctx), Domain(ctx))
			}
		}
	}
}

func TestNewPolicyRejects(t *testing.T) {
	cases := map[string]struct {
		keys  map[string]string
		allow map[string][]string
	}{
		"short key":       {map[string]string{"a": "short"}, nil},
		"shared key":      {map[string]string{"a": keyA, "b": keyA}, nil},
		"empty name":      {map[string]string{" ": keyA}, nil},
		"allow needs key": {map[string]string{"a": keyA}, map[string][]string{"b": {"/s.S/M"}}},
		"bad method":      {map[string]string{"a": keyA}, map[string][]string{"a": {"s.S/M"}}},
	}
	for name, c := range cases {
		if _, err := NewPolicy(c.keys, c.allow); err == nil {
			t.Errorf("%s: accepted", name)
		}
	}
}

func TestKeysFromEnv(t *testing.T) {
	t.Setenv("X_KEYS", " attendance = "+keyA+" ,kyc="+keyK+",")
	got, err := KeysFromEnv("X_KEYS")
	if err != nil || got["attendance"] != keyA || got["kyc"] != keyK || len(got) != 2 {
		t.Errorf("got %v %v", got, err)
	}
	for _, bad := range []string{"", "attendance", "=k", "a=" + keyA + ",a=" + keyK} {
		t.Setenv("X_KEYS", bad)
		if _, err := KeysFromEnv("X_KEYS"); err == nil {
			t.Errorf("%q: accepted", bad)
		}
	}
}

func TestInterceptorsEndToEnd(t *testing.T) {
	p := testPolicy(t)
	var serverCtx context.Context
	server := p.UnaryServerInterceptor()

	// The client interceptor's outgoing metadata becomes the server's
	// incoming metadata.
	invoker := func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		md, _ := metadata.FromOutgoingContext(ctx)
		_, err := server(metadata.NewIncomingContext(context.Background(), md), req,
			&grpc.UnaryServerInfo{FullMethod: method},
			func(ctx context.Context, req any) (any, error) { serverCtx = ctx; return "ok", nil })
		return err
	}
	client := UnaryClientInterceptor(keyA)
	if err := client(WithDomain(context.Background(), " Demo.Pineswift.com "), "/compat.Compat/ListPendingTasks", nil, nil, nil, invoker); err != nil {
		t.Fatal(err)
	}
	if Caller(serverCtx) != "attendance" || Domain(serverCtx) != "demo.pineswift.com" {
		t.Errorf("server saw caller %q domain %q", Caller(serverCtx), Domain(serverCtx))
	}
	if err := client(WithDomain(context.Background(), "demo.pineswift.com"), "/compat.Compat/GetProjectHours", nil, nil, nil, invoker); status.Code(err) != codes.PermissionDenied {
		t.Errorf("not allowed method: %v", err)
	}
}
