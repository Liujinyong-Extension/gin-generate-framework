package controllers

import (
	"fmt"
	"gin-generate-framework/app/validates"

	"github.com/gin-gonic/gin"
)

type TestController struct {
	BaseController
}
type TestIndexRequest struct {
	IndexRequest
}

func (test TestController) Index(c *gin.Context) {
	var request TestIndexRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		fmt.Println(err)
		test.ErrorJson(c, ParamError, err.Error())
		return
	}
	fmt.Println(request)

	if errors := validates.ValidateStruct(&request); errors != nil {
		fmt.Println(errors)
		for k, v := range errors {
			test.ErrorJson(c, ParamError, k+": "+v)
			return
		}
	}

	fmt.Println(request)
	test.SuccessJson(c, SuccessCode, "success", request)
}
