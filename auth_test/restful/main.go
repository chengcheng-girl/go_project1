package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.StaticFS("/public", http.Dir("C:/go_projects/demo/static"))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dst := path.Join("./static/upload", file.Filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dst)
		c.JSON(200, gin.H{
			"message": file.Filename+" uploaded!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
