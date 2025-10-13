package logic

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse 定义统一的错误响应结构
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse 创建新的错误响应
func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

// SendError 发送错误响应
func SendError(c *gin.Context, code int, message string) {
	c.JSON(code, NewErrorResponse(code, message))
}

// HandleBadRequest 处理请求参数错误
func HandleBadRequest(c *gin.Context, err error) {
	SendError(c, http.StatusBadRequest, err.Error())
}

// HandleNotFound 处理资源未找到错误
func HandleNotFound(c *gin.Context, message string) {
	SendError(c, http.StatusNotFound, message)
}

// HandleInternalError 处理内部服务器错误
func HandleInternalError(c *gin.Context, err error) {
	SendError(c, http.StatusInternalServerError, err.Error())
}
