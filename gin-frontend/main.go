package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	version = "dev"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":   "Gin Frontend",
			"version": version,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
