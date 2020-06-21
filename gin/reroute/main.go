package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setRoute() *gin.Engine {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "aha8080",
	//	})
	//})
	r.GET("/ping", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "0.0.0.0:7788")

	})
	return r
}

func main() {
	r := setRoute()
	r.Run(":8080")
}
