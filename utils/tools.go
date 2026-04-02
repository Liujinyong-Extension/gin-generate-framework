package utils

import (
	"gin-generate-framework/core/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logs(data map[string]interface{}, level logrus.Level, msg string, c *gin.Context) {
	global.Logrus.WithContext(c.Request.Context()).WithFields(data).Log(level, msg)
}
