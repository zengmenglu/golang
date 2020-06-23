package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func globalMid1(c *gin.Context) {
	fmt.Println("global-1-A")
	c.Next()
	fmt.Println("global-1-B")
}
func globalMid2(c *gin.Context) {
	fmt.Println("global-2-A")
	c.Next()
	fmt.Println("global-2-B")
}
func mid1(c *gin.Context) {
	fmt.Println("mid-1-A")
	c.Next()
	fmt.Println("mid-1-B")
}

func mid2(c *gin.Context) {
	fmt.Println("mid-2-A")
	c.Next()
	fmt.Println("mid-2-B")
}

func main() {
	r := gin.Default()
	rGrp := r.Group("", globalMid1)
	rGrp.Use(globalMid2)
	{
		rGrp.GET("/abc", mid1, mid2)
		rGrp.GET("/def", mid2)
	}
	r.Run(":8080")
}
