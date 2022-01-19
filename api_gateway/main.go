package main

import (
	"cyberblog_go/api_gateway/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	r := e.Group("api/v1.0")

	r.GET("/", controller.SayHello)

	panic(e.Run(":3000"))
}
