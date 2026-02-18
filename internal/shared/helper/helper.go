package helper

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseUintParam(c *gin.Context, name string) (uint, error) {
	v := c.Param(name)
	n, err := strconv.ParseUint(v, 10, 64)
	return uint(n), err
}
