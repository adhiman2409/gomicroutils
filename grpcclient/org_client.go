package grpcclient

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	"github.com/adhiman2409/gomicroutils/genproto/org"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	// read ca's cert
	caCert, err := os.ReadFile("/app/cert/ca-cert.pem")
	if err != nil {
		log.Fatal(caCert)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("/app/cert/client-cert.pem", "/app/cert/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	// create client connection
	conn, err := grpc.Dial(
		"organization-srv:50051",
		grpc.WithTransportCredentials(tlsCredential),
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
