package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin.Default
func setRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

func main() {
	r := setRoute()
	r.Run(":8080")
}
