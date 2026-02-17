package service

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"my-api/internal/modules/auth/models"
	"my-api/internal/modules/auth/repository"
	"my-api/internal/shared/errors"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repo      repository.UserRepositoryInterface
	jwtSecret []byte
}

func NewAuthService(repo repository.UserRepositoryInterface, secret string) *AuthService {
	return &AuthService{
		repo:      repo,
		jwtSecret: []byte(secret),
	}
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

	token, err := generateToken(user.ID, s.jwtSecret)
	if err != nil {
		return nil, nil, err
	}

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

	token, err := generateToken(user.ID, s.jwtSecret)
	if err != nil {
		return nil, nil, err
	}

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

func generateToken(userID uint, jwtSecret []byte) (*models.Token, error) {
	accessClaims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	access, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	return &models.Token{
		AccessToken: access,
		ExpiresIn:   3600,
	}, nil
}

func (s *AuthService) ExtractUserIDFromToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("token is invalid or expired")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("invalid token subject")
	}

	return uint(sub), nil
}
