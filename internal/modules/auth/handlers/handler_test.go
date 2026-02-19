package handlers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"my-api/internal/modules/auth/models"
	mocks "my-api/internal/modules/auth/repository"
	"my-api/internal/modules/auth/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const mockJWTSecret = "test"

func setupTestHandler() *AuthHandler {
	gin.SetMode(gin.TestMode)
	mockRepo := mocks.NewMockUserRepository()
	authService := service.NewAuthService(mockRepo, mockJWTSecret)
	return NewAuthHandler(authService)
}

func TestAuthHandler_Register(t *testing.T) {
	handler := setupTestHandler()
	router := gin.New()
	router.POST("/api/auth/register", handler.Register)

	req := models.RegisterRequest{
		Email:     "test@example.com",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
	}

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, httpReq)

	assert.Equal(t, 201, w.Code)

	var response models.RegisterResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotNil(t, response.User)
	assert.Equal(t, "test@example.com", response.User.Email)
	assert.NotNil(t, response.Token)
}

func TestAuthHandler_Login(t *testing.T) {
	mockRepo := mocks.NewMockUserRepository()
	hashedPass := service.NewAuthService(mockRepo, mockJWTSecret)

	// Register user first
	_, _, _ = hashedPass.Register("test@example.com", "password123", "John", "Doe")

	handler := NewAuthHandler(hashedPass)
	router := gin.New()
	router.POST("/api/auth/login", handler.Login)

	req := models.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, httpReq)

	assert.Equal(t, 200, w.Code)

	var response models.LoginResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.NotNil(t, response.User)
	assert.Equal(t, "test@example.com", response.User.Email)
	assert.NotNil(t, response.Token)
}

func TestAuthHandler_LoginInvalidCredentials(t *testing.T) {
	handler := setupTestHandler()
	router := gin.New()
	router.POST("/api/auth/login", handler.Login)

	req := models.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "password123",
	}

	body, _ := json.Marshal(req)
	httpReq := httptest.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, httpReq)

	assert.Equal(t, 401, w.Code)
}
