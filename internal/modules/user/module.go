package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-api/internal/modules/user/handlers"
	"my-api/internal/modules/user/repository"
	"my-api/internal/modules/user/service"
)

type Module struct {
	Handler    *handlers.UserHandler
	Service    *service.UserService
	Repository *repository.ProfileRepository
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewProfileRepository(db)
	svc := service.NewUserService(repo)
	handler := handlers.NewUserHandler(svc)

	return &Module{
		Handler:    handler,
		Service:    svc,
		Repository: repo,
	}
}

func (m *Module) RegisterRoutes(router *gin.Engine) {
	RegisterRoutes(router, m.Handler)
}
