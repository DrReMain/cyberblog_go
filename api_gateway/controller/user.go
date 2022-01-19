package controller

import (
	"context"
	pb "cyberblog_go/proto/hello"
	"fmt"
	"github.com/gin-gonic/gin"
	consulApi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func getAddr(name string) (string, error) {
	client, err := consulApi.NewClient(consulApi.DefaultConfig())
	if err != nil {
		return "", err
	}

	_, list, err := client.Agent().AgentHealthServiceByName(name)
	if err != nil {
		log.Printf("do not find: %s\n", name)
		return "", err
	}
	return fmt.Sprintf("%s:%d", list[0].Service.Address, list[0].Service.Port), nil
}

func SayHello(ctx *gin.Context) {
	addr, _ := getAddr("server_hello")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	res, err := client.SayHello(context.TODO(), &pb.HelloRequest{
		Name: "xiaobaozi",
	})
	if err != nil {
		log.Fatalln("call fail: ", err)
	}
	fmt.Println(res.Message)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": res.Message,
	})
}
