package api

import (
	"fmt"
	"gin-generate-framework/app/controllers"
	"gin-generate-framework/core/middleware"

	"github.com/gin-gonic/gin"
)

type TestApi struct {
}

func (TestApi TestApi) InitTestApi(r *gin.Engine) {
	fmt.Println("TestApi")
	group := r.Group("test")
	group.Use(middleware.AuthMiddleware())
	{
		group.GET("/index", controllers.TestController{}.Index)
		group.POST("/add", controllers.TestController{}.Add)
		group.PUT("/update", controllers.TestController{}.Update)
		group.DELETE("/delete", controllers.TestController{}.Delete)
	}
}
