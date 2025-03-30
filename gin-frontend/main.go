package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Gin Frontend",
		})
	})
	// 使用环境变量设置端口，默认值为 8080
	port := ":8080"
	r.Run(port)
}
