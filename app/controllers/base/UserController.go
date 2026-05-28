package base

import (
	"fmt"
	"gin-generate-framework/app/controllers"
	"gin-generate-framework/app/request"
	"gin-generate-framework/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.BaseController
}

func (UserController UserController) Login(c *gin.Context) {
	var LoginReq request.UserLoginRequest
	UserController.CheckInput(c, &LoginReq)

	fmt.Println(LoginReq, 2)
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

	//UserController.SuccessJson(c, controllers.SuccessCode, "success", one)
}
