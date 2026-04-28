package models

type Test struct {
	BaseModel
	Title    string `gorm:"type:varchar(255);comment:名称"`
	Content  string `gorm:"type:text;comment:内容"`
	Score    int    `gorm:"default:0;comment:分数"`
	Category string `gorm:"type:enum('apple','samsang','oppo');not null;default:apple;comment:分类 苹果:apple  三星:samsang 步步高:oppo"`
}

func (test Test) TableName() string {
	return "test"
}
