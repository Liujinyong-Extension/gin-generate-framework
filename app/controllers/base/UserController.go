package base

import (
	"gin-generate-framework/app/controllers"
	"gin-generate-framework/app/request"
	"gin-generate-framework/app/services"
	"gin-generate-framework/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.BaseController
}

func (UserController UserController) Login(c *gin.Context) {
	var LoginReq request.UserLoginRequest
	UserController.CheckInput(c, &LoginReq)
	userService := services.UserService{}
	one, err := userService.GetOne(request.WhereRequest{
		Conditions: []request.QueryCondition{
			{Field: "username", Operator: "eq", Value: LoginReq.UserName},
		},
	})

	if err != nil {
		UserController.ErrorJson(c, controllers.ServerErrorCode, err.Error())
		return
	}
	if one == nil {
		UserController.ErrorJson(c, controllers.ErrorCode, "用户不存在")
		return
	}
	userMap, ok := one.(map[string]interface{})
	if !ok {
		UserController.ErrorJson(c, controllers.ServerErrorCode, "数据类型错误")
		return
	}
	if userMap["password"] != utils.MD5(LoginReq.Password) {
		UserController.ErrorJson(c, controllers.ErrorCode, "密码错误")
		return
	}

	UserController.SuccessJson(c, controllers.SuccessCode, "登录成功", userMap)
}
