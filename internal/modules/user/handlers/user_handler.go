package handlers

import (
	"github.com/gin-gonic/gin"
	"my-api/internal/modules/user/models"
	"my-api/internal/modules/user/service"
	"my-api/internal/shared/errors"
	"strconv"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user_id"})
		return
	}

	profile, err := h.service.GetProfile(uint(userID))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, profile)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user_id"})
		return
	}

	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	profile, err := h.service.UpdateProfile(uint(userID), &req)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(200, profile)
}

func (h *UserHandler) CreateProfile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user_id"})
		return
	}

	profile, err := h.service.CreateProfile(uint(userID))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(201, profile)
}
