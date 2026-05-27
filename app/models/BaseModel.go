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

func (BaseModel BaseModel) Add(table string, data map[string]interface{}) (int64, error) {
	db := global.GormDB.Table(table)
	tx := db.Create(data)
	if tx.Error != nil {
		return 0, tx.Error
	}
	// 从 map 中获取 GORM 回填的自增 ID
	if id, ok := data["id"]; ok {
		switch v := id.(type) {
		case int64:
			return v, nil
		case float64:
			return int64(v), nil
		case int:
			return int64(v), nil
		}
	}
	return tx.RowsAffected, nil
}
func (BaseModel BaseModel) Update(table string, data map[string]interface{}) (int64, error) {
	// 从 data 中提取 id，然后从更新数据中移除
	id, ok := data["id"]
	if !ok {
		return 0, fmt.Errorf("缺少id参数")
	}
	delete(data, "id")

	db := global.GormDB.Table(table)

	// 判断记录是否存在
	var count int64
	if err := db.Where("id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, fmt.Errorf("记录不存在")
	}

	// 执行更新
	tx := db.Session(&gorm.Session{}).Where("id = ?", id).Updates(data)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
}
func (BaseModel BaseModel) Delete(table string, id int64) (int64, error) {
	db := global.GormDB.Table(table)
	var count int64
	if err := db.Where("id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, fmt.Errorf("记录不存在")
	}
	tx := db.Delete(nil)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return tx.RowsAffected, nil
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

	return total, result, err
}
func (BaseModel BaseModel) GetListNoPage(table string, conditions []request.QueryCondition) (int64, []interface{}, error) {
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
	err = db.Find(&list).Error

	// Convert to []interface{}
	result := make([]interface{}, len(list))
	for i, item := range list {
		result[i] = item
	}

	return total, result, err
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
