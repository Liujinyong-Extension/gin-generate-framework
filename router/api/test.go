package api

import (
	"fmt"
	"gin-generate-framework/app/controllers"

	"github.com/gin-gonic/gin"
)

type TestApi struct {
}

func (TestApi TestApi) InitTestApi(r *gin.Engine) {
	fmt.Println("TestApi")
	group := r.Group("test")
	{
		group.GET("/index", controllers.TestController{}.Index)
		group.POST("/add", controllers.TestController{}.Add)
		group.PUT("/update", controllers.TestController{}.Update)
		group.DELETE("/delete", controllers.TestController{}.Delete)
	}
}
