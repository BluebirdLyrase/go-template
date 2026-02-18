package address

import (
	"my-api/internal/modules/option/address/handlers"
	"my-api/internal/modules/option/address/repository"
	"my-api/internal/modules/option/address/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Module struct {
	Handler    *handlers.AddressHandler
	Service    *service.AddressService
	Repository *repository.AddressRepository
	jwtSecret  []byte
}

func NewModule(db *gorm.DB, secret string) *Module {
	repo := repository.NewAddressRepository(db)
	svc := service.NewAddressService(repo)
	handler := handlers.NewAddressHandler(svc)

	return &Module{
		Handler:    handler,
		Service:    svc,
		Repository: repo,
		jwtSecret:  []byte(secret),
	}
}

func (m *Module) RegisterRoutes(router *gin.RouterGroup) {
	RegisterRoutes(router, m.Handler)
}
