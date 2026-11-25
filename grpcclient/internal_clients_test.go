package grpcclient

import (
	"context"
	"net"
	"strings"
	"testing"

	"github.com/adhiman2409/gomicroutils/genproto/taskmgrcompat"
	"github.com/adhiman2409/gomicroutils/grpcauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

func TestDialInternalNeedsKey(t *testing.T) {
	t.Setenv("X_CALLER_KEY", "")
	if _, err := dialInternal("X_URL", "x:1", "X_CALLER_KEY"); err == nil || !strings.Contains(err.Error(), "X_CALLER_KEY") {
		t.Errorf("missing key: %v", err)
	}
}

type compatServer struct {
	taskmgrcompat.UnimplementedCompatServer
}

func (compatServer) GetProjectHours(ctx context.Context, req *taskmgrcompat.GetProjectHoursRequest) (*taskmgrcompat.GetProjectHoursResponse, error) {
	out := &taskmgrcompat.GetProjectHoursResponse{}
	for _, id := range req.ProjectIds {
		out.Projects = append(out.Projects, &taskmgrcompat.ProjectHours{ProjectId: grpcauth.Domain(ctx) + "/" + id, ActualHours: 1})
	}
	return out, nil
}

// TestCompatRoundTrip runs a generated Compat server behind grpcauth over an
// in-memory connection, with the same client interceptor dialInternal uses.
func TestCompatRoundTrip(t *testing.T) {
	kycKey := strings.Repeat("k", grpcauth.MinKeyLength)
	policy, err := grpcauth.NewPolicy(
		map[string]string{"kyc": kycKey},
		map[string][]string{"kyc": {"/taskmgrcompat.Compat/GetProjectHours"}})
	if err != nil {
		t.Fatal(err)
	}
	lis := bufconn.Listen(1 << 20)
	srv := grpc.NewServer(grpc.UnaryInterceptor(policy.UnaryServerInterceptor()))
	taskmgrcompat.RegisterCompatServer(srv, compatServer{})
	go srv.Serve(lis)
	t.Cleanup(srv.Stop)

	dial := func(key string) taskmgrcompat.CompatClient {
		conn, err := grpc.NewClient("passthrough:///bufnet",
			grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return lis.Dial() }),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(grpcauth.UnaryClientInterceptor(key)))
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() { conn.Close() })
		return taskmgrcompat.NewCompatClient(conn)
	}

	ctx := grpcauth.WithDomain(context.Background(), "demo.unirms.com")
	res, err := dial(kycKey).GetProjectHours(ctx, &taskmgrcompat.GetProjectHoursRequest{ProjectIds: []string{"P1"}})
	if err != nil || len(res.Projects) != 1 || res.Projects[0].ProjectId != "demo.unirms.com/P1" {
		t.Fatalf("allowed call: %v %v", res, err)
	}
	if _, err := dial(kycKey).ListOpenTasks(ctx, &taskmgrcompat.ListTasksRequest{}); status.Code(err) != codes.PermissionDenied {
		t.Errorf("method not allowed: %v", err)
	}
	if _, err := dial(strings.Repeat("x", grpcauth.MinKeyLength)).GetProjectHours(ctx, &taskmgrcompat.GetProjectHoursRequest{}); status.Code(err) != codes.Unauthenticated {
		t.Errorf("wrong key: %v", err)
	}
}
