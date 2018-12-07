package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/vulpemventures/bitcoin-grpc/proto"
	server "github.com/vulpemventures/bitcoin-grpc/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 5000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Running at localhost:", *port)
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	serveInstance := &server.BitcoinServer{}
	err = serveInstance.NewServer("localhost:19001", "admin1", "123")
	if err != nil {
		log.Fatalf("Failed to start bitcoin %v", err)
	}
	pb.RegisterBitcoinServer(grpcServer, serveInstance)
	grpcServer.Serve(lis)
}
