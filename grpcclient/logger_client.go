package grpcclient

import (
	"log"

	"github.com/adhiman2409/gomicroutils/genproto/logger"
	"google.golang.org/grpc"
)

type LoggerClient struct {
	conn   *grpc.ClientConn
	client logger.LoggerServiceClient
}

var loggerClient *LoggerClient

func GetLoggerClient() *LoggerClient {
	return loggerClient
}

func StartLoggerClient() {

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// create client connection
	conn, err := grpc.Dial(
		"logger-srv:50051",
		grpc.WithTransportCredentials(tlsCredentials),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := logger.NewLoggerServiceClient(conn)

	loggerClient = &LoggerClient{
		conn:   conn,
		client: client,
	}
}

func StopLoggerClient() {
	orgClient.conn.Close()
}
