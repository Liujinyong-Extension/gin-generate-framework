package services

import "gin-generate-framework/app/models"

type BaseService struct {
}

func (BaseService BaseService) GetList(model models.BaseModel, page int, page_size int) (int64, []models.BaseModel, error) {
	var list []models.BaseModel

	return int64(len(list)), list, nil
}
