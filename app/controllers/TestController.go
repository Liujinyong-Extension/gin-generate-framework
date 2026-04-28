package controllers

import (
	"fmt"
	"gin-generate-framework/app/models"
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
	//global.Redis.Set(c.Request.Context(), "test", "test", 60*time.Second)
	//time.Sleep(60 * time.Second)

	utils.Logs(map[string]interface{}{
		"page_num":  request.PageNum,
		"page_size": request.PageSize,
	}, logrus.InfoLevel, "这是一个测试", c)

	total, list, err := models.Test{}.GetList("test", request.PageNum, request.PageSize)
	if err != nil {
		test.ErrorJson(c, ServerErrorCode, err.Error())
		return
	}

	test.SuccessJson(c, SuccessCode, "success", map[string]interface{}{
		"total": total,
		"list":  list,
	})
}
