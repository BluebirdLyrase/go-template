package handlers

import (
	"github.com/gin-gonic/gin"
	"my-api/internal/modules/product/models"
	"my-api/internal/modules/product/service"
	"my-api/internal/shared/errors"
	"strconv"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.CreateProduct(&req)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(201, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	product, err := h.service.GetProduct(uint(id))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.UpdateProduct(uint(id), &req)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid product id"})
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(204, nil)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, products)
}
