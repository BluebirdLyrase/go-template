package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"my-api/internal/modules/auth/models"
	"my-api/internal/modules/auth/repository"
	"my-api/internal/shared/errors"
)

type AuthService struct {
	repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(email, password, firstName, lastName string) (*models.User, *models.Token, error) {
	// Check if user already exists
	existing, _ := s.repo.GetByEmail(email)
	if existing != nil {
		return nil, nil, errors.ErrUserAlreadyExists
	}

	// Create new user
	user := &models.User{
		Email:     email,
		Password:  hashPassword(password),
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}

	// Save user
	if err := s.repo.Create(user); err != nil {
		return nil, nil, errors.ErrInternalServer
	}

	// Generate token
	token := generateToken(user.ID)

	return user, token, nil
}

func (s *AuthService) Login(email, password string) (*models.User, *models.Token, error) {
	// Find user by email
	user, err := s.repo.GetByEmail(email)
	if err != nil || user == nil {
		return nil, nil, errors.ErrInvalidCredentials
	}

	// Verify password
	if !verifyPassword(user.Password, password) {
		return nil, nil, errors.ErrInvalidCredentials
	}

	// Generate token
	token := generateToken(user.ID)

	return user, token, nil
}

func (s *AuthService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil || user == nil {
		return nil, errors.ErrUserNotFound
	}
	return user, nil
}

// Helper functions
func hashPassword(password string) string {
	h := md5.New()
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func verifyPassword(hash, password string) bool {
	return hash == hashPassword(password)
}

func generateToken(userID uint) *models.Token {
	return &models.Token{
		AccessToken:  fmt.Sprintf("token_%d_%d", userID, time.Now().Unix()),
		RefreshToken: fmt.Sprintf("refresh_%d_%d", userID, time.Now().Unix()),
		ExpiresIn:    3600,
	}
}
