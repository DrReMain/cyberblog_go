package handler

import (
	"context"
	pb "cyberblog_go/proto/user"
	"cyberblog_go/server_hello/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedUserServer
}

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Passwd   string `gorm:"not null"`
}

func (s *Server) RegisterByPasswd(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Request: %v\n", in)

	user := User{
		Username: in.Username,
		Passwd:   in.Passwd,
	}

	DB, _ := db.ConnectDatabase()
	DB.Table("table_users").Create(&user)

	return &pb.RegisterResponse{
		Msg: "ok",
	}, nil
}
