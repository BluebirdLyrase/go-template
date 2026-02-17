package handlers

import (
	"github.com/gin-gonic/gin"
	"my-api/internal/modules/auth/models"
	"my-api/internal/modules/auth/service"
	"my-api/internal/shared/errors"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.service.Register(req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(201, models.RegisterResponse{
		User:  user,
		Token: token,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, models.LoginResponse{
		User:  user,
		Token: token,
	})
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		errors.HandleError(c, errors.ErrUnauthorized)
		return
	}

	id := userID.(uint)
	user, err := h.service.GetUserByID(id)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, user)
}
