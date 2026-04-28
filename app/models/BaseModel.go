package models

import "time"

import "gin-generate-framework/core/global"

type BaseModel struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// todo 研究一下sql怎么写
func (BaseModel BaseModel) GetList(table string, page int, page_size int) (int, []interface{}, error) {
	var list []map[string]interface{}
	var total int64

	err := global.GormDB.Table(table).Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	// Get paginated data
	offset := (page - 1) * page_size
	err = global.GormDB.Table(table).Offset(offset).Limit(page_size).Find(&list).Error

	// Convert to []interface{}
	result := make([]interface{}, len(list))
	for i, item := range list {
		result[i] = item
	}

	return int(total), result, err
}
