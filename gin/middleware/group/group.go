package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func middleware1Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func middlewareAuth(c *gin.Context) {
	fmt.Println("do auth")
}

func simpleRoute() *gin.Engine {
	r := gin.New()
	r.GET("/ping", middleware1Ping)
	return r
}

func authRoute() *gin.Engine {
	r := gin.New()

	// way 1
	authGrp := r.Group("/way1")
	authGrp.Use(middlewareAuth)
	{
		authGrp.GET("/ping", middleware1Ping)
	}

	// way 2
	authGrp2 := r.Group("/way2", middlewareAuth)
	{
		authGrp2.GET("/ping", middleware1Ping)
	}
	return r

}

func main() {
	//r := simpleRoute()
	r := authRoute()
	r.Run(":8080")
}
