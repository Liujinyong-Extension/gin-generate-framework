package controllers

import (
	"math"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

const (
	// SuccessCode 成功状态码
	SuccessCode = 200
	// ParamError 参数错误状态码
	ParamError = 422
	// ErrorCode 客户端错误状态码
	ErrorCode = 400
	// ServerErrorCode 服务器错误状态码
	ServerErrorCode = 500
)

type ReturnSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ErrorSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (b *BaseController) SuccessJson(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, ReturnSuccess{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
func (b *BaseController) ErrorJson(c *gin.Context, code int, message string) {
	c.JSON(200, ErrorSuccess{
		Code:    code,
		Message: message,
	})
}

func (b *BaseController) ListSuccessJson(c *gin.Context, code int, message string, data interface{}, total int64, page_num int, page_size int) {
	b.SuccessJson(c, code, message, map[string]interface{}{
		"total_page": math.Ceil(float64(total) / float64(page_size)),
		"list":       data,
		"page_num":   page_num,
		"page_size":  page_size,
	})
}
