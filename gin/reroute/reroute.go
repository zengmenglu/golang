package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func reRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/pong", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// inner
	r.GET("/inner", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/pong")
	})

	// outter
	r.GET("/outer", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	// forward
	r.GET("/forward", func(c *gin.Context) {
		rsp, err := http.Get("https://www.baidu.com/")
		if err != nil || rsp == nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.DataFromReader(rsp.StatusCode, rsp.ContentLength, rsp.Header.Get("Content-Type"), rsp.Body, nil)
	})
	return r
}

func main() {
	r := reRoute()
	r.Run(":8080")
}
