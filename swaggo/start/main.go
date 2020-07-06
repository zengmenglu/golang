package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"

	_ "swaggo/start/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Swagger Get Started API
// @version 1.0
// @description This is a sample for getting started.

// @tag.name App Tag.
// @tag.description This is App tag
// @tag.docs.url http://swagger.io/terms/
// @tag.docs.description This is outer tag

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8089
// @basePath /

// @query.collection.format multi
// @schemes http https
// @x-example-key {"key": "value"}
func main(){
	r := gin.Default()
	config := &ginSwagger.Config{
		URL: "http://localhost:8089/swagger/doc.json", //The url pointing to API definition
	}
	r.GET("/ping", get)
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config,swaggerFiles.Handler))
	r.Run(":8089")
}

// @description 进行路由检测
// @id example-get-id
// @summary 用于路由检测
// @tags example
// @accept json
// @produce  json
// @param id query int false "user id"
// @success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @failure 400 {string} json "{"code":400,"data":{},"msg":"fail"}"
// @router /ping [get]
func get(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}