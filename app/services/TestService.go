package services

import (
	"gin-generate-framework/app/models"
	"gin-generate-framework/app/request"
)

type TestService struct {
	BaseService
}

func (TestService TestService) GetList(request request.IndexRequest) (int64, []interface{}, error) {

	var list []interface{}

	total, list, err := models.Test{}.GetList(models.Test{}.TableName(), request.PageNum, request.PageSize)
	if err != nil {
		return 0, nil, err
	}
	return int64(total), list, nil
}
