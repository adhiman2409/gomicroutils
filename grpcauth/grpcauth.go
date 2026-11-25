// Package grpcauth checks which internal service is calling a gRPC server
// and which tenant the call is for.
//
// Each calling service has a secret caller key. The server holds a Policy:
// the keys it accepts, which caller each key belongs to, and the methods each
// caller may call. Every call also names its tenant domain in metadata; the
// server takes the domain from there (Domain(ctx)) and never from a request
// field, so a caller cannot act for a tenant by editing a message body.
//
// Server:
//
//	keys, err := grpcauth.KeysFromEnv("TASKMGR_GRPC_CALLER_KEYS") // "attendance=…,employee=…"
//	policy, err := grpcauth.NewPolicy(keys, map[string][]string{
//		"attendance": {"/taskmgrcompat.Compat/ListPendingTasks", "/taskmgrcompat.Compat/AddApprovedHours"},
//		"kyc":        {"/taskmgrcompat.Compat/GetProjectHours"},
//	})
//	srv := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(policy.UnaryServerInterceptor()))
//
// Client:
//
//	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(creds),
//		grpc.WithUnaryInterceptor(grpcauth.UnaryClientInterceptor(os.Getenv("TASKMGR_CALLER_KEY"))))
//	res, err := client.GetProjectHours(grpcauth.WithDomain(ctx, domain), req)
package grpcauth

import (
	"context"
	"crypto/subtle"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Metadata keys. gRPC metadata keys are lower case.
const (
	CallerKeyHeader = "x-caller-key"
	DomainHeader    = "x-tenant-domain"
)

// MinKeyLength is the shortest caller key NewPolicy accepts.
const MinKeyLength = 32

// Policy is the set of callers a server accepts and what each may call.
type Policy struct {
	callers []caller
	allow   map[string][]string
}

type caller struct {
	name string
	key  []byte
}

// NewPolicy builds a policy from caller keys (caller name -> key) and an
// allowlist (caller name -> full method names). A method entry is either a
// full method, "/package.Service/Method", or a whole service,
// "/package.Service/*". Every caller in the allowlist needs a key, and keys
// must be distinct and at least MinKeyLength long.
func NewPolicy(keys map[string]string, allow map[string][]string) (*Policy, error) {
	p := &Policy{allow: map[string][]string{}}
	seen := map[string]string{}
	names := make([]string, 0, len(keys))
	for name := range keys {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		key := keys[name]
		if strings.TrimSpace(name) == "" {
			return nil, errors.New("grpcauth: empty caller name")
		}
		if len(key) < MinKeyLength {
			return nil, fmt.Errorf("grpcauth: key for %q is shorter than %d characters", name, MinKeyLength)
		}
		if other, dup := seen[key]; dup {
			return nil, fmt.Errorf("grpcauth: %q and %q share a key", other, name)
		}
		seen[key] = name
		p.callers = append(p.callers, caller{name: name, key: []byte(key)})
	}
	for name, methods := range allow {
		if _, ok := keys[name]; !ok {
			return nil, fmt.Errorf("grpcauth: caller %q is in the allowlist but has no key", name)
		}
		for _, m := range methods {
			if !strings.HasPrefix(m, "/") || strings.Count(m, "/") != 2 {
				return nil, fmt.Errorf("grpcauth: %q is not a full method name like /package.Service/Method", m)
			}
		}
		p.allow[name] = append([]string(nil), methods...)
	}
	return p, nil
}

// KeysFromEnv reads caller keys from the named environment variable, written
// as comma-separated caller=key pairs.
func KeysFromEnv(name string) (map[string]string, error) {
	raw := strings.TrimSpace(os.Getenv(name))
	if raw == "" {
		return nil, fmt.Errorf("grpcauth: %s is not set", name)
	}
	out := map[string]string{}
	for _, pair := range strings.Split(raw, ",") {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		caller, key, ok := strings.Cut(pair, "=")
		caller, key = strings.TrimSpace(caller), strings.TrimSpace(key)
		if !ok || caller == "" || key == "" {
			return nil, fmt.Errorf("grpcauth: %s has an entry that is not caller=key", name)
		}
		if _, dup := out[caller]; dup {
			return nil, fmt.Errorf("grpcauth: %s lists %q twice", name, caller)
		}
		out[caller] = key
	}
	return out, nil
}

// callerFor returns the caller whose key matches, comparing every key in
// constant time so the answer's timing does not reveal which keys exist.
func (p *Policy) callerFor(key string) (string, bool) {
	found := ""
	for _, c := range p.callers {
		if subtle.ConstantTimeCompare([]byte(key), c.key) == 1 {
			found = c.name
		}
	}
	return found, found != ""
}

func (p *Policy) allowed(callerName, fullMethod string) bool {
	svc := fullMethod
	if i := strings.LastIndex(fullMethod, "/"); i > 0 {
		svc = fullMethod[:i] + "/*"
	}
	for _, m := range p.allow[callerName] {
		if m == fullMethod || m == svc {
			return true
		}
	}
	return false
}

type ctxKey int

const (
	callerCtxKey ctxKey = iota
	domainCtxKey
)

// Check authorises one call. It returns the context the handler should use,
// carrying the caller name and tenant domain, or a gRPC status error:
// Unauthenticated for a missing or unknown key, PermissionDenied for a method
// the caller may not call, InvalidArgument for a missing domain.
func (p *Policy) Check(ctx context.Context, fullMethod string) (context.Context, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	key := first(md, CallerKeyHeader)
	if key == "" {
		return nil, status.Error(codes.Unauthenticated, "missing caller key")
	}
	name, ok := p.callerFor(key)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unknown caller key")
	}
	if !p.allowed(name, fullMethod) {
		return nil, status.Errorf(codes.PermissionDenied, "%s may not call %s", name, fullMethod)
	}
	domain := strings.ToLower(strings.TrimSpace(first(md, DomainHeader)))
	if domain == "" {
		return nil, status.Error(codes.InvalidArgument, "missing tenant domain")
	}
	ctx = context.WithValue(ctx, callerCtxKey, name)
	return context.WithValue(ctx, domainCtxKey, domain), nil
}

// UnaryServerInterceptor applies Check to every unary call.
func (p *Policy) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		ctx, err := p.Check(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// Caller returns the caller name Check put in the context, or "".
func Caller(ctx context.Context) string {
	s, _ := ctx.Value(callerCtxKey).(string)
	return s
}

// Domain returns the tenant domain Check put in the context, or "".
func Domain(ctx context.Context) string {
	s, _ := ctx.Value(domainCtxKey).(string)
	return s
}

// WithDomain returns a context whose outgoing calls name the tenant domain.
func WithDomain(ctx context.Context, domain string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, DomainHeader, strings.ToLower(strings.TrimSpace(domain)))
}

// UnaryClientInterceptor adds the caller key to every outgoing unary call.
// The tenant domain is set per call with WithDomain.
func UnaryClientInterceptor(key string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(metadata.AppendToOutgoingContext(ctx, CallerKeyHeader, key), method, req, reply, cc, opts...)
	}
}

func first(md metadata.MD, k string) string {
	if v := md.Get(k); len(v) > 0 {
		return v[0]
	}
	return ""
}
