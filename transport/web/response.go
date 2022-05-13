package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type error1 struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Errorf(message)
	c.AbortWithStatusJSON(statusCode, error1{message})
}
