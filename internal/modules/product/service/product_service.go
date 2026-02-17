package service

import (
	"my-api/internal/modules/product/models"
	"my-api/internal/modules/product/repository"
	"my-api/internal/shared/errors"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(req *models.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		IsActive:    true,
	}

	if err := s.repo.Create(product); err != nil {
		return nil, errors.ErrInternalServer
	}

	return product, nil
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil || product == nil {
		return nil, errors.ErrProductNotFound
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(id uint, req *models.UpdateProductRequest) (*models.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil || product == nil {
		return nil, errors.ErrProductNotFound
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}
	if req.Category != "" {
		product.Category = req.Category
	}

	if err := s.repo.Update(product); err != nil {
		return nil, errors.ErrInternalServer
	}

	return product, nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}

func (s *ProductService) GetAllProducts() ([]*models.Product, error) {
	return s.repo.GetAll()
}
