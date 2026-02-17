package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrUserNotFound        = NewError(404, "user not found")
	ErrUserAlreadyExists   = NewError(400, "user already exists")
	ErrInvalidCredentials  = NewError(401, "invalid credentials")
	ErrUnauthorized        = NewError(401, "unauthorized")
	ErrProductNotFound     = NewError(404, "product not found")
	ErrInvalidEmail        = NewError(400, "invalid email format")
	ErrInvalidPassword     = NewError(400, "password must be at least 6 characters")
	ErrInternalServer      = NewError(500, "internal server error")
)

type AppError struct {
	Code    int
	Message string
}

func NewError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (e *AppError) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Code, ErrorResponse{Code: appErr.Code, Message: appErr.Message})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:    500,
		Message: "Internal server error",
	})
}
