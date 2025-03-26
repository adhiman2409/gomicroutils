package grpcclient

import (
	"fmt"
	"log"

	"github.com/adhiman2409/gomicroutils/genproto/notification"
	"google.golang.org/grpc"
)

type NotificationClient struct {
	conn   *grpc.ClientConn
	client notification.NotificationServiceClient
}

var notificationClient *NotificationClient

func GetNotificationClient() *NotificationClient {
	return notificationClient
}

func StartNotificationClient() {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	conn, err := grpc.NewClient(
		"notification-srv:50051",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		fmt.Println("Error in dialing:", err)
		log.Fatal(err)
	}

	client := notification.NewNotificationServiceClient(conn)

	notificationClient = &NotificationClient{
		conn:   conn,
		client: client,
	}
}

func StopNotificationClient() {
	orgClient.conn.Close()
}
