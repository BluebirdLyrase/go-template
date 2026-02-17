package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"my-api/internal/modules/auth/models"
	"my-api/testdata/mocks"
)

func TestAuthService_Register(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	service := NewAuthService(mockRepo)

	user, token, err := service.Register("test@example.com", "password123", "John", "Doe")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, token)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "John", user.FirstName)
	assert.NotEmpty(t, token.AccessToken)
}

func TestAuthService_RegisterDuplicate(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	mockRepo.Users = append(mockRepo.Users, &models.User{
		ID:    1,
		Email: "test@example.com",
	})
	service := NewAuthService(mockRepo)

	_, _, err := service.Register("test@example.com", "password123", "John", "Doe")

	assert.Error(t, err)
	assert.Equal(t, "user already exists", err.Error())
}

func TestAuthService_Login(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	hashedPass := hashPassword("password123")
	mockRepo.Users = append(mockRepo.Users, &models.User{
		ID:        1,
		Email:     "test@example.com",
		Password:  hashedPass,
		FirstName: "John",
	})

	service := NewAuthService(mockRepo)

	user, token, err := service.Login("test@example.com", "password123")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, token)
	assert.Equal(t, "test@example.com", user.Email)
	assert.NotEmpty(t, token.AccessToken)
}

func TestAuthService_LoginInvalidPassword(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	mockRepo.Users = append(mockRepo.Users, &models.User{
		ID:       1,
		Email:    "test@example.com",
		Password: hashPassword("password123"),
	})

	service := NewAuthService(mockRepo)

	_, _, err := service.Login("test@example.com", "wrongpassword")

	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
}

func TestAuthService_LoginUserNotFound(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	service := NewAuthService(mockRepo)

	_, _, err := service.Login("nonexistent@example.com", "password123")

	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
}

func TestAuthService_GetUserByID(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	mockRepo.Users = append(mockRepo.Users, &models.User{
		ID:        1,
		Email:     "test@example.com",
		FirstName: "John",
	})

	service := NewAuthService(mockRepo)

	user, err := service.GetUserByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestAuthService_GetUserByIDNotFound(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	service := NewAuthService(mockRepo)

	_, err := service.GetUserByID(999)

	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}
