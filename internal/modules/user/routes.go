package user

import (
	"github.com/gin-gonic/gin"
	"my-api/internal/modules/user/handlers"
)

func RegisterRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	users := router.Group("/api/users")
	{
		users.GET("/:user_id/profile", handler.GetProfile)
		users.PUT("/:user_id/profile", handler.UpdateProfile)
		users.POST("/:user_id/profile", handler.CreateProfile)
	}
}
