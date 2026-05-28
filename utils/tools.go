package utils

import (
	"crypto/md5"
	"encoding/hex"
	"gin-generate-framework/core/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logs(data map[string]interface{}, level logrus.Level, msg string, c *gin.Context) {
	global.Logrus.WithContext(c.Request.Context()).WithFields(data).Log(level, msg)
}

// MD5 计算字符串的 MD5 哈希值
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
