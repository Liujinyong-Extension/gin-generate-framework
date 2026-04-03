package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// 数据库连接实例
	GormDB *gorm.DB

	// 验证器实例，用于结构体验证
	Validate *validator.Validate

	// 业务日志记录器，用于记录业务逻辑日志
	Logrus *logrus.Logger

	// 访问日志记录器，用于记录HTTP访问日志
	AccessLog *logrus.Logger

	// Redis实例
	Redis *redis.Client
)
