package controllers

import (
	"fmt"
	"gin-generate-framework/app/validates"
	"math"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	User map[string]interface{}
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
	// AuthorErrorCode 认证错误状态码
	AuthorErrorCode = 401
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
	c.AbortWithStatusJSON(200, ErrorSuccess{
		Code:    code,
		Message: message,
	})
}

func (b *BaseController) ListSuccessJson(c *gin.Context, code int, message string, data interface{}, total int64, page_num int, page_size int) {
	var totalPage float64
	if page_size <= 0 {
		totalPage = 0
	} else {
		totalPage = math.Ceil(float64(total) / float64(page_size))
	}
	b.SuccessJson(c, code, message, map[string]interface{}{
		"total_page": totalPage,
		"list":       data,
		"page_num":   page_num,
		"page_size":  page_size,
	})
}

// GetLoginedUser 从请求上下文中获取已登录用户信息（由 AuthMiddleware 设置）
func (b *BaseController) GetLoginedUser(c *gin.Context) map[string]interface{} {
	if user, exists := c.Get("user"); exists {
		if userMap, ok := user.(map[string]interface{}); ok {
			b.User = userMap
			return userMap
		}
	}
	return nil
}

func (b *BaseController) CheckInput(c *gin.Context, req interface{}) {
	method := c.Request.Method

	var err error

	if method == "GET" {
		err = c.ShouldBindQuery(req)
	} else {
		err = c.ShouldBindJSON(req)
	}

	if err != nil {
		fmt.Println(err, 1)
		b.ErrorJson(c, ParamError, err.Error())
		return
	}

	if errors := validates.ValidateStruct(req); errors != nil {
		fmt.Println(errors, 2)
		for k, v := range errors {
			b.ErrorJson(c, ParamError, k+": "+v)
			return
		}
	}

}
