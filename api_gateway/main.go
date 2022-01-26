package main

import (
	"cyberblog_go/api_gateway/user"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	e := gin.Default()

	r := e.Group("api/v1.0")

	r.POST("/register/passwd", user.RegisterByPasswd)

	log.Panic(e.Run(":3000"))
}
