package grpcclient

import (
	"log"

	"github.com/adhiman2409/gomicroutils/genproto/org"
	"google.golang.org/grpc"
)

type OrgClient struct {
	conn   *grpc.ClientConn
	client org.OrgServiceClient
}

var orgClient *OrgClient

func GetOrgClient() *OrgClient {
	return orgClient
}

func StartOrgClient() {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	conn, err := grpc.Dial(
		"organization-srv:50051",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := org.NewOrgServiceClient(conn)

	orgClient = &OrgClient{
		conn:   conn,
		client: client,
	}
}

func StopOrgClient() {
	orgClient.conn.Close()
}
