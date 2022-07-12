package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFound api response of not found
func NotFound(c *gin.Context, what string) {
	msg := "not found"
	if len(what) > 0 {
		msg = fmt.Sprintf("%s not found", what)
	}
	c.JSON(http.StatusNotFound, gin.H{
		"code": http.StatusNotFound,
		"msg":  msg,
	})
}

// Timeout api response of timeout
func Timeout(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{
		"code": http.StatusRequestTimeout,
		"msg":  "request timeout",
	})
}

// OK build ok body
func OK(c *gin.Context, payload interface{}) {
	if payload == nil {
		c.JSON(http.StatusOK, gin.H{"code": 0})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"payload": payload,
		})
	}
}
