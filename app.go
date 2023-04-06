package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func fallbackString(value, fallback string) string {
	if (len(value) == 0) {
		return fallback
	}

	return value
}

func main() {
	router := gin.Default()
	port := fallbackString(os.Getenv("PORT"), "8080")

	router.GET("/*path", func(context *gin.Context) {
		path := context.Param("path")
		context.JSON(200, gin.H{
			"path": path,
		})
	})

	router.Run(":" + port)
}