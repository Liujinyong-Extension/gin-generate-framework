package services

import "gin-generate-framework/app/request"

// Model 约束模型必须实现的方法
type Model interface {
	TableName() string
	GetList(table string, page int, pageSize int) (int64, []interface{}, error)
}

// BaseService 泛型基础服务，M 为具体的模型类型
type BaseService[M Model] struct {
}

// GetList 通用分页列表查询
func (BaseService[M]) GetList(req request.IndexRequest) (int64, []interface{}, error) {
	var m M
	total, list, err := m.GetList(m.TableName(), req.PageNum, req.PageSize)
	if err != nil {
		return 0, nil, err
	}
	return total, list, nil
}
