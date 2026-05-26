package models

import (
	"fmt"
	"time"

	"gin-generate-framework/app/request"
	"gin-generate-framework/core/global"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (BaseModel BaseModel) GetList(table string, page int, page_size int, conditions []request.QueryCondition) (int64, []interface{}, error) {
	var list []map[string]interface{}
	var total int64

	db := global.GormDB.Table(table)

	// 应用自定义查询条件
	db = applyConditions(db, conditions)

	err := db.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	// Get paginated data
	offset := (page - 1) * page_size
	err = db.Offset(offset).Limit(page_size).Find(&list).Error

	// Convert to []interface{}
	result := make([]interface{}, len(list))
	for i, item := range list {
		result[i] = item
	}

	return int64(total), result, err
}

// applyConditions 将条件应用到 GORM 查询中
// 支持的操作符：like, eq, ne, gt, gte, lt, lte, in
func applyConditions(db *gorm.DB, conditions []request.QueryCondition) *gorm.DB {
	for _, cond := range conditions {
		switch cond.Operator {
		case "like":
			db = db.Where(fmt.Sprintf("`%s` LIKE ?", cond.Field), fmt.Sprintf("%%%v%%", cond.Value))
		case "eq":
			db = db.Where(fmt.Sprintf("`%s` = ?", cond.Field), cond.Value)
		case "ne":
			db = db.Where(fmt.Sprintf("`%s` != ?", cond.Field), cond.Value)
		case "gt":
			db = db.Where(fmt.Sprintf("`%s` > ?", cond.Field), cond.Value)
		case "gte":
			db = db.Where(fmt.Sprintf("`%s` >= ?", cond.Field), cond.Value)
		case "lt":
			db = db.Where(fmt.Sprintf("`%s` < ?", cond.Field), cond.Value)
		case "lte":
			db = db.Where(fmt.Sprintf("`%s` <= ?", cond.Field), cond.Value)
		case "in":
			db = db.Where(fmt.Sprintf("`%s` IN (?)", cond.Field), cond.Value)
		}
	}
	return db
}
