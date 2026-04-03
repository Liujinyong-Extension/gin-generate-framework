package core

import (
	"context"
	"fmt"
	"gin-generate-framework/core/global"
	"io"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() {

	InitViper()
	InitDatabase()
	InitRedis()
	InitValidate()
	InitLog()
}
func InitViper() {
	env := func() string {
		if e := os.Getenv("ENV"); e != "" {
			return e
		}
		return "local"
	}()
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
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
	global.Logrus = logrus.New()
	global.Logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	global.Logrus.SetLevel(logrus.InfoLevel)
	global.Logrus.SetOutput(getLogWriter("logs/business.log"))
	global.Logrus.AddHook(&ContextHook{}) // 自定义 Hook（可选，见下文）

	global.AccessLog = logrus.New()
	global.AccessLog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	global.AccessLog.SetOutput(getLogWriter("logs/access.log"))
	global.AccessLog.SetLevel(logrus.InfoLevel)

}

// 日志切割（按天）
func getLogWriter(filePath string) io.Writer {
	writer, _ := rotatelogs.New(
		filePath+".%Y%m%d",
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	return writer
}

type ContextHook struct{}

func (hook *ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *ContextHook) Fire(entry *logrus.Entry) error {
	// 从 context 中提取 trace_id（需要 Gin 中间件注入到 request.Context）
	if entry.Context != nil {
		if traceID, ok := entry.Context.Value("trace_id").(string); ok {
			entry.Data["trace_id"] = traceID
		}
	}
	return nil
}
func InitRedis() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := global.Redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis connect errer", err)

	}
}
