package handlers

import (
	"my-api/internal/modules/auth/models"
	"my-api/internal/modules/auth/service"
	"my-api/internal/shared/errors"
	"strings"

	"github.com/gin-gonic/gin"
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
	auth := c.GetHeader("Authorization")
	if auth == "" {
		errors.HandleError(c, errors.ErrUnauthorized, "missing Authorization header")
		return
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 {
		errors.HandleError(c, errors.ErrUnauthorized, "invalid Authorization header format")
		return
	}

	userID, err := h.service.ExtractUserIDFromToken(parts[1])
	if err != nil {
		errors.HandleError(c, errors.ErrUnauthorized, err.Error())
		return
	}

	user, err := h.service.GetUserByID(userID)
	if err != nil {
		errors.HandleError(c, errors.ErrNotFound, err.Error())
		return
	}

	c.JSON(200, user)
}
