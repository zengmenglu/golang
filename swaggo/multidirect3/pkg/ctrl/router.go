package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"swaggo/multidirect3/pkg/ctrl/v0"
	_ "swaggo/multidirect3/docs"
)

func HttpRouter() *gin.Engine {
	router := gin.Default()
	servGroup := router.Group("api")
	{
		servGroup.GET("/get", v0.Get)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
