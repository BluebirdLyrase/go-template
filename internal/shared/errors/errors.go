package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

var (
	ErrUserNotFound       = NewError(404, "user not found")
	ErrUserAlreadyExists  = NewError(400, "user already exists")
	ErrInvalidCredentials = NewError(401, "invalid credentials")
	ErrUnauthorized       = NewError(401, "unauthorized")
	ErrProductNotFound    = NewError(404, "product not found")
	ErrInvalidEmail       = NewError(400, "invalid email format")
	ErrInvalidPassword    = NewError(400, "password must be at least 6 characters")
	ErrInternalServer     = NewError(500, "internal server error")
)

type AppError struct {
	Code    int
	Message string
}

type SimpleError struct {
	Message string
}

func (e *SimpleError) Error() string {
	return e.Message
}

func New(message string) error {
	return &SimpleError{Message: message}
}

func NewError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func (e *AppError) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, err error, detail ...string) {
	var d string

	if len(detail) > 0 && detail[0] != "" {
		d = detail[0]
	} else {
		d = err.Error()
	}

	if appErr, ok := err.(*AppError); ok {
		c.JSON(appErr.Code, ErrorResponse{
			Code:    appErr.Code,
			Message: appErr.Message,
			Detail:  d,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:    500,
		Message: "Internal server error",
		Detail:  d,
	})
}
