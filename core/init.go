package core

import (
	"fmt"
	"gin-generate-framework/core/global"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() {
	InitViper()
	InitDatabase()
	InitValidate()
	InitLog()
}
func InitViper() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	fmt.Println("当前的环境是:", env)
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("core/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
}
func InitDatabase() {
	Username := viper.GetString("database.user")
	Password := viper.GetString("database.password")
	Host := viper.GetString("database.host")
	PORT := viper.GetString("database.port")
	DBNAME := viper.GetString("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, PORT, DBNAME)

	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	global.GormDB, err = db, err
	if err != nil {
		fmt.Println("mysql connect errer", err)
	}
	if global.GormDB.Error != nil {
		fmt.Println("database connect errer", global.GormDB.Error)
	}
	//todo 这个地方回头可以加在配置文件里
	sqlDB, _ := global.GormDB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println(global.GormDB)
}

func InitValidate() {
	global.Validate = validator.New()
	global.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// 优先获取 json tag，其次获取 form tag
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return "" // 忽略该字段
		}
		if name == "" {
			name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]
			if name == "-" {
				return "" // 忽略该字段
			}
		}
		if name != "" {
			return name
		}
		// 如果没有 json 和 form tag，返回字段原名
		return fld.Name
	})
}
func InitLog() {
	logrus.Debug("this is logrus debug")
}
