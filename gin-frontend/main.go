package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	version = "dev"
)

func main() {
	logger := log.New(os.Stdout, "GIN: ", log.LstdFlags|log.Lshortfile)

	r := gin.Default()
	r.Use(gin.LoggerWithWriter(logger.Writer()))

	if err := r.LoadHTMLGlob("templates/*"); err != nil {
		logger.Fatalf("Failed to load templates: %v", err)
	}

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

	logger.Printf("Starting server on port %s", port)
	if err := r.Run(":" + port); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
