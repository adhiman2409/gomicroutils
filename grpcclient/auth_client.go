package grpcclient

import (
	"log"

	"github.com/adhiman2409/gomicroutils/genproto/auth"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	conn   *grpc.ClientConn
	client auth.AuthServiceClient
}

var grpcClient *GrpcClient

func GetAuthClient() *GrpcClient {
	return grpcClient
}

func StopGrpcClient() {
	grpcClient.conn.Close()
}

func StartAuthClient() {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	conn, err := grpc.Dial(
		"auth-srv:50051",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := auth.NewAuthServiceClient(conn)

	grpcClient = &GrpcClient{
		conn:   conn,
		client: client,
	}
}
