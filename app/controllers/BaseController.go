package controllers

import "github.com/gin-gonic/gin"

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

type IndexRequest struct {
	PageNum  int `form:"page_num" validate:"required,number,min=1"`
	PageSize int `form:"page_size" validate:"required,number,min=1"`
}
type detailRequest struct {
	Id int `form:"id" validate:"required,number"`
}

//TODO 想想怎么实现增删改查集合到一起  ai给的建议是用反射
