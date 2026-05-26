package request

type IndexRequest struct {
	PageRequest
	Conditions []QueryCondition // 自定义查询条件，由 ParseConditions 解析填充
}
