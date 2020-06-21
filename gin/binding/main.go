package main

import "github.com/gin-gonic/gin"

func setRoute() *gin.Engine {
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	return router
}

func main() {
	router := setRoute()
	router.Run(":8080")
}
