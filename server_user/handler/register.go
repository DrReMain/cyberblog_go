package handler

import (
	"context"
	pb "cyberblog_go/proto/user"
	"log"
)

type Server struct {
	pb.UnimplementedUserServer
}

func (s *Server) RegisterByPasswd(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Request: %v\n", in)

	// TODO
	return &pb.RegisterResponse{
		Msg: "ok",
	}, nil
}
