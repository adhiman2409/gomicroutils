package grpcclient

// Clients for the internal gRPC servers added for Task Manager v1. Unlike the
// older clients, each one checks in with a caller key (package grpcauth), its
// address can be overridden from the environment, Start returns an error
// instead of exiting, and the generated client is embedded as is, so callers
// use the genproto request and response types directly.
//
// Every call must name the tenant: wrap the context with
// grpcauth.WithDomain(ctx, domain).
//
//	if err := grpcclient.StartTaskMgrCompatClient(); err != nil { log.Fatal(err) }
//	res, err := grpcclient.GetTaskMgrCompatClient().GetProjectHours(
//		grpcauth.WithDomain(ctx, domain), &taskmgrcompat.GetProjectHoursRequest{ProjectIds: ids})

import (
	"fmt"
	"os"

	"github.com/adhiman2409/gomicroutils/genproto/availability"
	"github.com/adhiman2409/gomicroutils/genproto/helpdeskticket"
	"github.com/adhiman2409/gomicroutils/genproto/projectdirectory"
	"github.com/adhiman2409/gomicroutils/genproto/taskmgrcompat"
	"github.com/adhiman2409/gomicroutils/grpcauth"
	"google.golang.org/grpc"
)

// dialInternal opens a TLS connection that sends the caller key from keyEnv
// on every call. The address comes from addrEnv, or def when it is unset.
func dialInternal(addrEnv, def, keyEnv string) (*grpc.ClientConn, error) {
	key := os.Getenv(keyEnv)
	if key == "" {
		return nil, fmt.Errorf("%s is not set", keyEnv)
	}
	addr := os.Getenv(addrEnv)
	if addr == "" {
		addr = def
	}
	creds, err := loadTLSCredentials()
	if err != nil {
		return nil, fmt.Errorf("load TLS credentials: %w", err)
	}
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(grpcauth.UnaryClientInterceptor(key)))
	if err != nil {
		return nil, fmt.Errorf("dial %s: %w", addr, err)
	}
	return conn, nil
}

// ---- ProjectDirectory (kyc) ----

type ProjectDirectoryClient struct {
	conn *grpc.ClientConn
	projectdirectory.ProjectDirectoryClient
}

var projectDirectoryClient *ProjectDirectoryClient

// StartProjectDirectoryClient connects to kyc's ProjectDirectory.
// Env: PROJECTDIRECTORY_SERVICE_URL (default kyc-srv:50051),
// PROJECTDIRECTORY_CALLER_KEY (required).
func StartProjectDirectoryClient() error {
	conn, err := dialInternal("PROJECTDIRECTORY_SERVICE_URL", "kyc-srv:50051", "PROJECTDIRECTORY_CALLER_KEY")
	if err != nil {
		return fmt.Errorf("project directory client: %w", err)
	}
	projectDirectoryClient = &ProjectDirectoryClient{conn, projectdirectory.NewProjectDirectoryClient(conn)}
	return nil
}

// GetProjectDirectoryClient returns the client, or nil before Start.
func GetProjectDirectoryClient() *ProjectDirectoryClient { return projectDirectoryClient }

func StopProjectDirectoryClient() {
	if projectDirectoryClient != nil {
		projectDirectoryClient.conn.Close()
	}
}

// ---- Availability (attendance) ----

type AvailabilityClient struct {
	conn *grpc.ClientConn
	availability.AvailabilityClient
}

var availabilityClient *AvailabilityClient

// StartAvailabilityClient connects to attendance's Availability.
// Env: AVAILABILITY_SERVICE_URL (default attendance-srv:50051),
// AVAILABILITY_CALLER_KEY (required).
func StartAvailabilityClient() error {
	conn, err := dialInternal("AVAILABILITY_SERVICE_URL", "attendance-srv:50051", "AVAILABILITY_CALLER_KEY")
	if err != nil {
		return fmt.Errorf("availability client: %w", err)
	}
	availabilityClient = &AvailabilityClient{conn, availability.NewAvailabilityClient(conn)}
	return nil
}

// GetAvailabilityClient returns the client, or nil before Start.
func GetAvailabilityClient() *AvailabilityClient { return availabilityClient }

func StopAvailabilityClient() {
	if availabilityClient != nil {
		availabilityClient.conn.Close()
	}
}

// ---- Compat (taskmgr) ----

type TaskMgrCompatClient struct {
	conn *grpc.ClientConn
	taskmgrcompat.CompatClient
}

var taskMgrCompatClient *TaskMgrCompatClient

// StartTaskMgrCompatClient connects to taskmgr's Compat.
// Env: TASKMGR_COMPAT_SERVICE_URL (default taskmgr-srv:50051),
// TASKMGR_COMPAT_CALLER_KEY (required).
func StartTaskMgrCompatClient() error {
	conn, err := dialInternal("TASKMGR_COMPAT_SERVICE_URL", "taskmgr-srv:50051", "TASKMGR_COMPAT_CALLER_KEY")
	if err != nil {
		return fmt.Errorf("taskmgr compat client: %w", err)
	}
	taskMgrCompatClient = &TaskMgrCompatClient{conn, taskmgrcompat.NewCompatClient(conn)}
	return nil
}

// GetTaskMgrCompatClient returns the client, or nil before Start.
func GetTaskMgrCompatClient() *TaskMgrCompatClient { return taskMgrCompatClient }

func StopTaskMgrCompatClient() {
	if taskMgrCompatClient != nil {
		taskMgrCompatClient.conn.Close()
	}
}

// ---- HelpdeskTicket (helpdesk) ----

type HelpdeskTicketClient struct {
	conn *grpc.ClientConn
	helpdeskticket.HelpdeskTicketClient
}

var helpdeskTicketClient *HelpdeskTicketClient

// StartHelpdeskTicketClient connects to helpdesk's HelpdeskTicket.
// Env: HELPDESK_TICKET_SERVICE_URL (default helpdesk-srv:50051),
// HELPDESK_TICKET_CALLER_KEY (required).
func StartHelpdeskTicketClient() error {
	conn, err := dialInternal("HELPDESK_TICKET_SERVICE_URL", "helpdesk-srv:50051", "HELPDESK_TICKET_CALLER_KEY")
	if err != nil {
		return fmt.Errorf("helpdesk ticket client: %w", err)
	}
	helpdeskTicketClient = &HelpdeskTicketClient{conn, helpdeskticket.NewHelpdeskTicketClient(conn)}
	return nil
}

// GetHelpdeskTicketClient returns the client, or nil before Start.
func GetHelpdeskTicketClient() *HelpdeskTicketClient { return helpdeskTicketClient }

func StopHelpdeskTicketClient() {
	if helpdeskTicketClient != nil {
		helpdeskTicketClient.conn.Close()
	}
}
