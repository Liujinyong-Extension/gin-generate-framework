package api

import (
	"gin-generate-framework/app/controllers"

	"github.com/gin-gonic/gin"
)

func TestApi(r *gin.Engine) *gin.Engine {
	//查詢
	r.GET("/test", controllers.TestController{}.Index)
	r.POST("/add", controllers.TestController{}.Add)
	r.PUT("/update", controllers.TestController{}.Update)
	r.DELETE("/delete", controllers.TestController{}.Delete)
	return r
}
