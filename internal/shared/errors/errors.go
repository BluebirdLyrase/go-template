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
	ErrNotFound           = NewError(404, "Not Found")
	ErrBadRequest         = NewError(400, "Bad Request")
	ErrInvalidCredentials = NewError(401, "Invalid Credentials")
	ErrUnauthorized       = NewError(401, "Unauthorized")
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
