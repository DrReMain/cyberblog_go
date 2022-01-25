package main

import (
	pb "cyberblog_go/proto/user"
	"cyberblog_go/server_hello/handler"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	IP   = "localhost"
	PORT = 8080
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	svc := grpc.NewServer()
	pb.RegisterUserServer(svc, &handler.Server{})
	log.Printf("server listening at %v\n", listen.Addr())
	if err := svc.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
