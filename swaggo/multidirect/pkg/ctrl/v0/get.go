package v0

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 路由检测
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /ping [get]
func Get(c *gin.Context){
	c.JSON(http.StatusOK,"get success")
}
