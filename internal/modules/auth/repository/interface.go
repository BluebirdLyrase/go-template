package repository

import "my-api/internal/modules/auth/models"

type UserRepositoryInterface interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}
