package auth

import (
	"my-api/internal/modules/auth/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *handlers.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
		auth.GET("/me", handler.GetMe)
	}
}
