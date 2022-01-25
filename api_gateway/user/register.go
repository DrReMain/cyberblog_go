package user

import (
	"context"
	pb "cyberblog_go/proto/user"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type RegisterRequestBody struct {
	Userame string `json:"username"`
	Passwd  string `json:"passwd"`
}

func RegisterByPasswd(ctx *gin.Context) {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: %v\n", err)
		return
	}
	defer conn.Close()

	json := RegisterRequestBody{}
	err = ctx.BindJSON(&json)
	if err != nil {
		log.Printf("bind json fail: %v\n", err)
		return
	}

	client := pb.NewUserClient(conn)
	res, err := client.RegisterByPasswd(context.TODO(), &pb.RegisterRequest{
		Username: json.Userame,
		Passwd:   json.Passwd,
	})
	if err != nil {
		log.Println("call rpc fail: ", err)
		return
	}

	// TODO
	log.Printf("Received From Server: %v\n", res)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": res.Msg,
	})
}
