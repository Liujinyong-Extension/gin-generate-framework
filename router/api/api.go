package api

import (
	"gin-generate-framework/app/controllers"

	"github.com/gin-gonic/gin"
)

func TestApi(r *gin.Engine) *gin.Engine {
	r.GET("/test", controllers.TestController{}.Index)
	return r
}
