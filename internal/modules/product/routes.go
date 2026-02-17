package product

import (
	"github.com/gin-gonic/gin"
	"my-api/internal/modules/product/handlers"
)

func RegisterRoutes(router *gin.Engine, handler *handlers.ProductHandler) {
	products := router.Group("/api/products")
	{
		products.GET("", handler.GetAllProducts)
		products.GET("/:id", handler.GetProduct)
		products.POST("", handler.CreateProduct)
		products.PUT("/:id", handler.UpdateProduct)
		products.DELETE("/:id", handler.DeleteProduct)
	}
}
