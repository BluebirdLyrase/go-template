package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return r
}
