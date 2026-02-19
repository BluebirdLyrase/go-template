package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupGinContext(paramKey, paramValue string) (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = gin.Params{
		{
			Key:   paramKey,
			Value: paramValue,
		},
	}

	return c, w
}

func TestParseUintParam_Success(t *testing.T) {
	c, _ := setupGinContext("id", "123")

	result, err := ParseUintParam(c, "id")

	assert.NoError(t, err)
	assert.Equal(t, uint(123), result)
}

func TestParseUintParam_InvalidNumber(t *testing.T) {
	c, _ := setupGinContext("id", "abc")

	result, err := ParseUintParam(c, "id")

	assert.Error(t, err)
	assert.Equal(t, uint(0), result)
}

func TestParseUintParam_EmptyParam(t *testing.T) {
	c, _ := setupGinContext("id", "")

	result, err := ParseUintParam(c, "id")

	assert.Error(t, err)
	assert.Equal(t, uint(0), result)
}
