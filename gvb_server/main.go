package main

import (
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/router"
)

func main() {
	// 配置文件的读取
	core.InitConfig()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	router := router.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Info("程序开始运行:", addr)
	router.Run(addr)
}
