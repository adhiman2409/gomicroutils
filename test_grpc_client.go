package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/adhiman2409/gomicroutils/genproto/monitor"
)

func main() {
	// Server address
	serverAddr := "grpc.unirms.com:443"

	// Create TLS credentials (for HTTPS/TLS connection)
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false, // Set to true only for testing with self-signed certs
	}
	creds := credentials.NewTLS(tlsConfig)

	// Connect to the gRPC server
	fmt.Printf("Connecting to gRPC server at %s...\n", serverAddr)
	conn, err := grpc.Dial(
		serverAddr,
		grpc.WithTransportCredentials(creds),
		grpc.WithBlock(),
		grpc.WithTimeout(10*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	fmt.Println("✓ Successfully connected to gRPC server!")

	// Create a client
	client := pb.NewMonitorServiceClient(conn)

	// Test 1: Heartbeat
	fmt.Println("\n--- Testing Heartbeat RPC ---")
	testHeartbeat(client)

	// Test 2: UpdateCheckInCheckoutStatus
	fmt.Println("\n--- Testing UpdateCheckInCheckoutStatus RPC ---")
	testCheckInCheckout(client)
}

func testHeartbeat(client pb.MonitorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.HeartbeatRequest{
		EmployeeId: "test-employee-123",    // Replace with a real employee ID
		Domain:     "example.com",          // Replace with your domain
	}

	fmt.Printf("Sending Heartbeat request: %+v\n", req)
	resp, err := client.Heartbeat(ctx, req)
	if err != nil {
		log.Printf("❌ Heartbeat failed: %v", err)
		return
	}

	fmt.Printf("✓ Heartbeat successful!\n")
	fmt.Printf("  Date: %d-%02d-%02d\n", resp.Year, resp.Month, resp.Day)
	fmt.Printf("  Checked In: %v\n", resp.IsCheckedIn)
	if resp.IsCheckedIn {
		fmt.Printf("  Check-in Time: %s (Source: %s)\n", resp.CheckinTimestamp, resp.CheckinSource)
	}
	fmt.Printf("  Checked Out: %v\n", resp.IsCheckedOut)
	if resp.IsCheckedOut {
		fmt.Printf("  Check-out Time: %s (Source: %s)\n", resp.CheckoutTimestamp, resp.CheckoutSource)
	}
	fmt.Printf("  On Leave: %v\n", resp.IsOnLeave)
	fmt.Printf("  Holiday: %v\n", resp.IsHoliday)
	fmt.Printf("  Weekly Off: %v\n", resp.IsWeeklyOffDay)
	fmt.Printf("  Message: %s\n", resp.Message)
}

func testCheckInCheckout(client pb.MonitorServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	req := &pb.EmployeeCheckInCheckOutRequest{
		Day:               int32(now.Day()),
		Month:             int32(now.Month()),
		Year:              int32(now.Year()),
		EmployeeId:        "test-employee-123",  // Replace with a real employee ID
		IsCheckedIn:       true,
		CheckInTimestamp:  now.Format(time.RFC3339),
		CheckInSource:     "grpc-test-client",
		IsCheckedOut:      false,
		CheckOutTimestamp: "",
		CheckOutSource:    "",
		Domain:            "example.com",  // Replace with your domain
	}

	fmt.Printf("Sending UpdateCheckInCheckoutStatus request: %+v\n", req)
	resp, err := client.UpdateCheckInCheckoutStatus(ctx, req)
	if err != nil {
		log.Printf("❌ UpdateCheckInCheckoutStatus failed: %v", err)
		return
	}

	fmt.Printf("✓ UpdateCheckInCheckoutStatus successful!\n")
	fmt.Printf("  Status: %v\n", resp.Status)
	fmt.Printf("  Message: %s\n", resp.Message)
}
