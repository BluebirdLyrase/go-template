package auth

import (
	"my-api/internal/modules/auth/handlers"
	"my-api/internal/modules/auth/repository"
	"my-api/internal/modules/auth/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Handler    *handlers.AuthHandler
	Service    *service.AuthService
	Repository *repository.UserRepository
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewUserRepository(db)
	svc := service.NewAuthService(repo)
	handler := handlers.NewAuthHandler(svc)

	return &Module{
		Handler:    handler,
		Service:    svc,
		Repository: repo,
	}
}

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	RegisterRoutes(router, m.Handler)
}
