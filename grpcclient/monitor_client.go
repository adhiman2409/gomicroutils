package grpcclient

import (
	"fmt"
	"log"

	"github.com/adhiman2409/gomicroutils/genproto/monitor"
	"google.golang.org/grpc"
)

type MonitorClient struct {
	conn   *grpc.ClientConn
	client monitor.MonitorServiceClient
}

var monitorClient *MonitorClient

func GetMonitorClient() *MonitorClient {
	return monitorClient
}

func StartMonitorClient() {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	conn, err := grpc.NewClient(
		"monitor-srv:50051",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		fmt.Println("Error in dialing:", err)
		log.Fatal(err)
	}

	client := monitor.NewMonitorServiceClient(conn)

	monitorClient = &MonitorClient{
		conn:   conn,
		client: client,
	}
}

func StopMMonitorClient() {
	orgClient.conn.Close()
}
