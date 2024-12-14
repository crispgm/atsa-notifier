// Package handler .
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CodeSuccess .
const CodeSuccess = 0

// Other error codes
const (
	CodeParamsErr = iota + 1001
	CodeLoadPlayer
	CodeLoadTemplate
)

// Status .
type Status struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

// Response .
type Response struct {
	Status

	Data interface{} `json:"data"`
}

// SuccessResponse build a default response with data
func SuccessResponse(c *gin.Context, data any) {
	resp := &Response{
		Status: Status{
			StatusCode:    CodeSuccess,
			StatusMessage: "success",
		},
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

// ErrorResponse build a default response with data
func ErrorResponse(c *gin.Context, code int, msg string, data any) {
	resp := &Response{
		Status: Status{
			StatusCode:    code,
			StatusMessage: msg,
		},
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func getLogger(c *gin.Context) *logrus.Entry {
	logger, ok := c.Get("logger")
	if !ok {
		logger = logrus.WithField("context", "missing")
	}
	logEntry := logger.(*logrus.Entry)

	return logEntry
}
