package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-api/internal/modules/product/handlers"
	"my-api/internal/modules/product/repository"
	"my-api/internal/modules/product/service"
)

type Module struct {
	Handler    *handlers.ProductHandler
	Service    *service.ProductService
	Repository *repository.ProductRepository
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)
	handler := handlers.NewProductHandler(svc)

	return &Module{
		Handler:    handler,
		Service:    svc,
		Repository: repo,
	}
}

func (m *Module) RegisterRoutes(router *gin.Engine) {
	RegisterRoutes(router, m.Handler)
}
