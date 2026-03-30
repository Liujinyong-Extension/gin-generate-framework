package router

import (
	"gin-generate-framework/router/api"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api.TestApi(r)
	return r
}
