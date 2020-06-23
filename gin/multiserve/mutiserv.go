package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

func init() {
	SetServers()
}

func route() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

// mutiple server
var (
	g        errgroup.Group
	portList = []string{":8081", ":8082", ":8083"}
	servers  = make(map[string]*http.Server, len(portList))
)

// multiple server definition
func SetServers() {
	fmt.Println("begin set proxy servers...")
	for _, port := range portList {
		servers[port] = &http.Server{
			Addr:         port,
			Handler:      route(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
	}
}

func main() {
	for port := range servers {
		serv := servers[port]
		g.Go(func() error {
			return serv.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil { // 阻塞进程直至所有goroutine返回
		log.Fatal(err)
	}
}
