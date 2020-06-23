package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// upload single file
// curl -X POST http://localhost:8080/upload -F "file=@C:\test\test.png" -H "Content-Type: multipart/form-data" -H "dst: C:\test2\test.png"
func uploadSingleRoute() *gin.Engine {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file") // 默认内存限制32M
		if err != nil {
			fmt.Printf("upload err:%s\n", err)
			c.String(http.StatusInternalServerError, fmt.Sprintf("uploaded fail,err:%s", err))
			return
		}
		fmt.Printf("file:%s\n", file.Filename)

		dst := c.Request.Header.Get("dst")
		fmt.Printf("dst:%s\n", dst)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Printf("save err:%s\n", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("save fail,err:%s", err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%s uploaded success", file.Filename))
	})
	return r
}

// upload multiple files
// curl -X POST http://localhost:8080/upload -F "file=@C:\test\pic1.png" -F "file=@C:\test\pic2.png" -H "Content-Type: multipart/form-data" -H "dst: C:\test2\\"
func uploadMultiRoute() *gin.Engine {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		dst := c.Request.Header.Get("dst")
		form, err := c.MultipartForm() // 默认内存限制32M
		if err != nil {
			fmt.Printf("upload err:%s\n", err)
			c.String(http.StatusInternalServerError, fmt.Sprintf("uploaded fail,err:%s", err))
			return
		}
		files := form.File["file"] // 读取标记为file的文件
		for _, file := range files {
			dstfile := fmt.Sprintf("%s%s", dst, file.Filename)
			fmt.Printf("file:%s,dst file:%s\n", file.Filename, dstfile)
			err = c.SaveUploadedFile(file, dstfile)
			if err != nil {
				fmt.Printf("save err:%s\n", err)
				c.String(http.StatusBadRequest, fmt.Sprintf("save fail,err:%s", err))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("uploaded success"))
	})
	return r
}

func main() {
	//r := uploadSingleRoute()
	r := uploadMultiRoute()
	r.Run(":8080")
}
