package main

import (
	"fmt"
	"gin-generate-framework/core"
	"gin-generate-framework/router"

	"github.com/spf13/viper"
)

func main() {
	//加载配置
	core.Init()
	//加载路由
	r := router.SetupRouter()

	serverPort := viper.GetInt("server.port")
	r.Run(":" + fmt.Sprintf("%d", serverPort))
}
