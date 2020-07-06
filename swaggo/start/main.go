package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"

	_ "swaggo/start/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8089
// @BasePath /
func main(){
	r := gin.Default()
	config := &ginSwagger.Config{
		URL: "http://localhost:8089/swagger/doc.json", //The url pointing to API definition
	}
	r.GET("/ping", get)
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config,swaggerFiles.Handler))
	r.Run(":8089")
}

// @Summary 路由检测
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /ping [get]
func get(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}