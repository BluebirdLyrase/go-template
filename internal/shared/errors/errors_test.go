package errors

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupGin() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func TestSimpleError_Error(t *testing.T) {
	err := New("test error")

	assert.Equal(t, "test error", err.Error())
}

func TestNewError(t *testing.T) {
	appErr := NewError(400, "bad request")

	assert.Equal(t, 400, appErr.Code)
	assert.Equal(t, "bad request", appErr.Message)
	assert.Equal(t, "bad request", appErr.Error())
}

func TestHandleError_WithAppError(t *testing.T) {
	c, w := setupGin()

	err := NewError(404, "Not Found")
	HandleError(c, err)

	assert.Equal(t, 404, w.Code)

	var resp ErrorResponse
	errDecode := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, errDecode)

	assert.Equal(t, 404, resp.Code)
	assert.Equal(t, "Not Found", resp.Message)
	assert.Equal(t, "Not Found", resp.Detail)
}

func TestHandleError_WithAppError_WithDetail(t *testing.T) {
	c, w := setupGin()

	err := NewError(401, "Unauthorized")
	HandleError(c, err, "custom detail")

	assert.Equal(t, 401, w.Code)

	var resp ErrorResponse
	errDecode := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, errDecode)

	assert.Equal(t, 401, resp.Code)
	assert.Equal(t, "Unauthorized", resp.Message)
	assert.Equal(t, "custom detail", resp.Detail)
}

func TestHandleError_WithGenericError(t *testing.T) {
	c, w := setupGin()

	err := New("some internal error")
	HandleError(c, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var resp ErrorResponse
	errDecode := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, errDecode)

	assert.Equal(t, 500, resp.Code)
	assert.Equal(t, "Internal server error", resp.Message)
	assert.Equal(t, "some internal error", resp.Detail)
}
