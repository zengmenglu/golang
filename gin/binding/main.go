package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Uri struct {
	User     string `uri:"user" binding:"required"`
	Password string `uri:"password" binding:"required"`
}

// this is for request uri param check
// uri: [GET] localhost:8080/myUser/myPassword
func setUriRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/:user/:password", func(c *gin.Context) {
		var uri Uri
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "request uri wrong",
				"error":  err.Error(),
			})
			return
		}
		fmt.Printf("uri:%+v\n", uri)
		if uri.User != "myUser" || uri.Password != "myPassword" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauth",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "log in",
		})

	})
	return router
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// this is for request data json param check
// uri: [POST] localhost:8080/login
// body: {"user": "myUser", "password": "myPassword"}
func setLoginRoute() *gin.Engine {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "request form wrong",
			})
			return
		}
		fmt.Printf("form:%+v\n", form)
		if form.User != "myUser" || form.Password != "myPassword" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauth",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "log in",
		})

	})
	return router
}

func main() {
	router := setUriRoute()
	//router := setLoginRoute()
	router.Run(":8080")
}
