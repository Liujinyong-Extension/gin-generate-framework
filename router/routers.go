package router

import (
	"gin-generate-framework/core/middleware"
	"gin-generate-framework/router/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//不加载中间件
	r := gin.New()
	//加载错误处理中间件
	r.Use(middleware.ErrorHandler())
	api.TestApi(r)
	return r
}
