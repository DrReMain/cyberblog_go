package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"t":       time.Now().UnixMilli(),
			"success": true,
			"result":  "HelloWorld",
			"code":    100000,
			"msg":     "OK",
		})
	})

	panic(r.Run(":8080"))
}
