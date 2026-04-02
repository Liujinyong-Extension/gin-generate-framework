package controllers

import (
	"fmt"
	"gin-generate-framework/app/validates"
	"gin-generate-framework/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	if errors := validates.ValidateStruct(&request); errors != nil {
		fmt.Println(errors)
		for k, v := range errors {
			test.ErrorJson(c, ParamError, k+": "+v)
			return
		}
	}
	utils.Logs(map[string]interface{}{
		"page_num":  request.PageNum,
		"page_size": request.PageSize,
	}, logrus.InfoLevel, "这是一个测试", c)

	test.SuccessJson(c, SuccessCode, "success", request)
}
