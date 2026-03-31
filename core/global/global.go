package global

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB

	Validate *validator.Validate
)
