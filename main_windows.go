//go:build windows
// +build windows

package main

import (
	"fmt"
	"gin-generate-framework/core"
	"gin-generate-framework/router"
	"runtime"

	"github.com/spf13/viper"
)

func main() {
	//加载配置
	core.Init()
	//加载路由
	r := router.SetupRouter()

	serverPort := viper.GetInt("server.port")
	addr := ":" + fmt.Sprintf("%d", serverPort)

	// Windows 环境使用普通启动
	fmt.Printf("当前操作系统: %s\n", runtime.GOOS)
	fmt.Println("使用普通启动模式")
	r.Run(addr)
}
