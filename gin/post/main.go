package main

import "github.com/gin-gonic/gin"

// this is for get param is post form
// uri: [POST]localhost:8080/form_post?id=123&page=1
// body:{"message": ahamessage, “name”: "nick"}
func setFormRoute() *gin.Engine {
	router := gin.Default()
	router.POST("/form_post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0") // if no data,return default value

		message := c.PostForm("message")
		name := c.DefaultPostForm("name", "anonymous") // if no data,return default value

		c.JSON(200, gin.H{
			"status":  "posted",
			"id":      id,
			"page":    page,
			"message": message,
			"name":    name,
		})
	})
	return router
}

func main() {
	router := setFormRoute()
	router.Run(":8080")
}
