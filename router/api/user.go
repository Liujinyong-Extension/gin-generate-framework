package api

import (
	"fmt"
	"gin-generate-framework/app/controllers/base"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func (UserApi UserApi) InitUserApi(r *gin.Engine) {
	fmt.Println("UserApi")
	group := r.Group("user")
	{
		group.POST("/login", base.UserController{}.Login) //登录接口
	}
}
