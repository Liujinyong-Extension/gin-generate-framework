package services

import "gin-generate-framework/app/request"

// Model 约束模型必须实现的方法
type Model interface {
	TableName() string
	GetList(table string, page int, pageSize int, conditions []request.QueryCondition) (int64, []interface{}, error)
	GetListNoPage(table string, conditions []request.QueryCondition) (int64, []interface{}, error)
	Add(table string, data map[string]interface{}) (int64, error)
	Update(table string, data map[string]interface{}) (int64, error)
	Delete(table string, id int64) (int64, error)
}

// BaseService 泛型基础服务，M 为具体的模型类型
type BaseService[M Model] struct {
}

// GetList 通用分页列表查询，支持自定义条件筛选
func (BaseService[M]) GetList(req request.PageRequest) (int64, []interface{}, error) {
	var m M
	total, list, err := m.GetList(m.TableName(), req.PageNum, req.PageSize, req.Conditions)
	if err != nil {
		return 0, nil, err
	}
	return total, list, nil
}
func (BaseService[M]) GetListNoPage(req request.PageRequest) (int64, []interface{}, error) {
	var m M
	total, list, err := m.GetListNoPage(m.TableName(), req.Conditions)
	if err != nil {
		return 0, nil, err
	}
	return total, list, nil
}
func (BaseService[M]) Add(req map[string]interface{}) (int64, error) {
	var m M
	return m.Add(m.TableName(), req)
}
func (BaseService[M]) Update(req map[string]interface{}) (int64, error) {
	var m M
	return m.Update(m.TableName(), req)
}
func (BaseService[M]) Delete(req request.IdRequest) (int64, error) {
	var m M
	return m.Delete(m.TableName(), int64(req.Id))
}
