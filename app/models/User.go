package models

type User struct {
	BaseModel
	Username string `gorm:"type:varchar(255);not null;comment:用户名"`
	Password string `gorm:"type:varchar(255);not null;comment:密码"`
}

func (user User) TableName() string {
	return "gin_user"
}
