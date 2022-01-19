package main

import (
	pb "cyberblog_go/proto/hello"
	"cyberblog_go/server_hello/handler"
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = 8080
	IP   = "192.168.50.143"
	NAME = "server_hello"
)

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", IP, PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	client, err := consulApi.NewClient(consulApi.DefaultConfig())
	if err != nil {
		log.Fatalf("consul client: %v\n", err)
	}

	reg := consulApi.AgentServiceRegistration{
		Tags:    []string{NAME},
		Name:    NAME,
		Address: IP,
		Port:    PORT,
		//Check: &consulApi.AgentServiceCheck{
		//	CheckID:                        "getimg service test",
		//	TCP:                            fmt.Sprintf("%s:%d", IP, PORT),
		//	Timeout:                        "5s",
		//	Interval:                       "5s",
		//	DeregisterCriticalServiceAfter: "30s",
		//},
	}

	err = client.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatalln("service register fail: ", err)
	}

	svc := grpc.NewServer()
	pb.RegisterGreeterServer(svc, &handler.Server{})
	log.Printf("server listening at %v\n", listen.Addr())
	if err := svc.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
