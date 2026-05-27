package request

import (
	"encoding/json"
	"net/url"
	"regexp"
)

/*
*
分页请求参数
*/
type PageRequest struct {
	PageNum    int              `form:"page_num" validate:"required,number,min=1"`
	PageSize   int              `form:"page_size" validate:"required,number,min=1"`
	Conditions []QueryCondition // 自定义查询条件，由 ParseConditions 解析填充
}

/*
*
id请求参数
*/
type IdRequest struct {
	Id int `json:"id" form:"id" validate:"required,number,min=1"`
}

// fieldNameRegex 校验字段名安全，防止 SQL 注入
var fieldNameRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

// QueryCondition 查询条件，从 URL 参数中解析
// 格式：?field=["operator","value"]
// 示例：?title=["like","测"]  →  WHERE title LIKE '%测%'
type QueryCondition struct {
	Field    string
	Operator string
	Value    interface{}
}

// ParseConditions 从 URL query 参数中解析自定义条件
// knownFields 为已知的业务字段（如 page_num、page_size），会被跳过不当作条件
func ParseConditions(query url.Values, knownFields ...string) []QueryCondition {
	known := make(map[string]bool, len(knownFields))
	for _, f := range knownFields {
		known[f] = true
	}

	var conditions []QueryCondition
	for key, values := range query {
		if known[key] {
			continue
		}
		if len(values) == 0 || values[0] == "" {
			continue
		}
		// 校验字段名安全
		if !fieldNameRegex.MatchString(key) {
			continue
		}

		// 尝试解析 JSON 数组 ["operator", "value"]
		var arr []interface{}
		if err := json.Unmarshal([]byte(values[0]), &arr); err == nil && len(arr) >= 2 {
			operator, ok := arr[0].(string)
			if ok && operator != "" {
				conditions = append(conditions, QueryCondition{
					Field:    key,
					Operator: operator,
					Value:    arr[1],
				})
			}
		}
	}
	return conditions
}
