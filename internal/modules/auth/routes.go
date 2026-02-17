package auth

import (
	"github.com/gin-gonic/gin"
	"my-api/internal/modules/auth/handlers"
)

func RegisterRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.GET("/me", handler.GetMe) // Protected endpoint
	}
}
