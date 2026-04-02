package router

import (
	"gin-generate-framework/core/middleware"
	"gin-generate-framework/router/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//创建一个不加载默认中间件的engine
	r := gin.New()
	//加载错误处理中间件
	r.Use(middleware.TraceMiddleware())
	r.Use(middleware.AccessLogMiddleware())
	r.Use(middleware.ErrorHandler())
	r.Use(gin.Recovery())
	api.TestApi(r)
	return r
}
