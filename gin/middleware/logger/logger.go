package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f) // 记录到文件。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // 同时将日志写入文件和控制台

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":8080")
}
