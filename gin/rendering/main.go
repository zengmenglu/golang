package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProtoData struct {
	Label *string `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Reps  []int64 `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
}

func renderRoute() *gin.Engine {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ // gin.H is a shortcut for map[string]interface{}
			"message": "json",
		})
	})

	r.GET("/morejson", func(c *gin.Context) {
		var jsonMsg struct {
			Name    string
			Message string
		}
		jsonMsg.Name = "Nick"
		jsonMsg.Message = "this is message"
		c.JSON(http.StatusOK, jsonMsg)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{ // gin.H is a shortcut for map[string]interface{}
			"message": "xml",
		})
	})

	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "yaml",
		})
	})

	return r
}

func main() {
	r := renderRoute()
	r.Run(":8080")
}
