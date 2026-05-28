package base

import (
	"gin-generate-framework/app/controllers"
	"gin-generate-framework/app/request"
	"gin-generate-framework/app/services"
	"gin-generate-framework/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
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

	// 创建 JWT 声明
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.MapClaims{
		"user_name": userMap["user_name"],
		"exp":       expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 从配置读取 JWT 密钥
	jwtKey := []byte(viper.GetString("jwt.secret"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		UserController.ErrorJson(c, controllers.ServerErrorCode, "生成Token失败")
		return
	}

	UserController.SuccessJson(c, controllers.SuccessCode, "登录成功", gin.H{"token": tokenString, "user": userMap})
}
