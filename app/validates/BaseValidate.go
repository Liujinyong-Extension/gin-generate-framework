package validates

import (
	"gin-generate-framework/core/global"
	"strings"

	"github.com/go-playground/validator/v10"
)

func init() {
	global.Validate = validator.New()

}

var errorMessages = map[string]string{
	"required": "字段为必填项",
	"email":    "邮箱格式不正确",
	"min":      "长度不能小于 %s",
	"max":      "长度不能大于 %s",
	"len":      "长度必须为 %s",
	"numeric":  "必须为数字",
	"url":      "URL格式不正确",
}

// 验证结构体并返回格式化的错误
func ValidateStruct(s interface{}) map[string]string {
	err := global.Validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		tag := err.Tag()
		param := err.Param()

		message := errorMessages[tag]
		if message == "" {
			message = tag
		}
		// 替换参数占位符
		message = strings.Replace(message, "%s", param, -1)

		errors[field] = message
	}
	return errors
}
