package handler

import (
	"context"
	pb "cyberblog_go/proto/hello"
	"fmt"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Printf("Received: %v\n", in.GetName())
	return &pb.HelloResponse{
		Message: "Hello " + in.GetName(),
	}, nil
}
