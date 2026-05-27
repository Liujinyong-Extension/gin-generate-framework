package controllers

import (
	"encoding/json"
	"fmt"
	"gin-generate-framework/app/request"
	"gin-generate-framework/app/services"
	"gin-generate-framework/app/validates"
	"gin-generate-framework/utils"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TestController struct {
	BaseController
}

func (test TestController) Index(c *gin.Context) {
	var requestParam request.PageRequest

	if err := c.ShouldBindQuery(&requestParam); err != nil {
		fmt.Println(err)
		test.ErrorJson(c, ParamError, err.Error())
		return
	}

	if errors := validates.ValidateStruct(&requestParam); errors != nil {
		fmt.Println(errors)
		for k, v := range errors {
			test.ErrorJson(c, ParamError, k+": "+v)
			return
		}
	}

	// 从 URL 中解析自定义查询条件（排除分页等已知参数）
	requestParam.Conditions = request.ParseConditions(c.Request.URL.Query(), "page_num", "page_size")

	utils.Logs(map[string]interface{}{
		"page_num":   requestParam.PageNum,
		"page_size":  requestParam.PageSize,
		"conditions": requestParam.Conditions,
	}, logrus.InfoLevel, "这是一个测试", c)

	total, list, err := services.TestService{}.GetList(requestParam)

	if err != nil {
		test.ErrorJson(c, ServerErrorCode, err.Error())
		return
	}
	test.ListSuccessJson(c, SuccessCode, "success", list, int64(math.Ceil(float64(total)/float64(requestParam.PageSize))), requestParam.PageNum, requestParam.PageSize)
}

func (test TestController) Add(c *gin.Context) {
	var requestParam request.TestAddRequest
	if err := c.ShouldBindJSON(&requestParam); err != nil {
		fmt.Println(err, 1)
		test.ErrorJson(c, ParamError, err.Error())
		return
	}
	if errors := validates.ValidateStruct(&requestParam); errors != nil {
		fmt.Println(errors, 2)
		for k, v := range errors {
			test.ErrorJson(c, ParamError, k+": "+v)
			return
		}
	}
	data, _ := json.Marshal(requestParam)
	var reqMap map[string]interface{}
	json.Unmarshal(data, &reqMap)
	now := time.Now().Format("2006-01-02 15:04:05")

	if requestParam.CreatedAt == nil {
		reqMap["created_at"] = now
	}
	if requestParam.UpdatedAt == nil {
		reqMap["updated_at"] = now
	}
	affected, err := services.TestService{}.Add(reqMap)

	if err != nil {
		test.ErrorJson(c, ServerErrorCode, err.Error())
		return
	}
	test.SuccessJson(c, SuccessCode, "success", map[string]interface{}{
		"affected": affected,
	})
}
func (test TestController) Update(c *gin.Context) {
	var updateReq request.TestUpdateRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		fmt.Println(err, 1)
		test.ErrorJson(c, ParamError, err.Error())
		return
	}
	if errors := validates.ValidateStruct(&updateReq); errors != nil {
		fmt.Println(errors, 2)
		for k, v := range errors {
			test.ErrorJson(c, ParamError, k+": "+v)
			return
		}
	}
	data, _ := json.Marshal(updateReq)
	var reqMap map[string]interface{}
	json.Unmarshal(data, &reqMap)
	now := time.Now().Format("2006-01-02 15:04:05")
	reqMap["updated_at"] = now
	affected, err := services.TestService{}.Update(reqMap)
	if err != nil {
		test.ErrorJson(c, ServerErrorCode, err.Error())
		return
	}
	test.SuccessJson(c, SuccessCode, "success", map[string]interface{}{
		"affected": affected,
	})
}
func (test TestController) Delete(c *gin.Context) {
	fmt.Println("delete")

	urlMap := map[string]string{
		"url1": "http://127.0.0.1:9090/update",
		"url2": "http://127.0.0.1:9090/update",
	}

	// 使用channel收集结果
	resChan := make(chan int, len(urlMap))
	var wg sync.WaitGroup

	// 并发发送请求
	for _, url := range urlMap {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resChan <- test.SendHttp(u)
		}(url)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// 收集结果
	resArr := []int{}
	for res := range resChan {
		resArr = append(resArr, res)
	}

	fmt.Printf("并发请求成功返回: %d\n", resArr)
}

func (test TestController) SendHttp(str string) int {
	// 发送PUT请求到/update端点
	req, err := http.NewRequest("PUT", str, nil)
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return 0
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
