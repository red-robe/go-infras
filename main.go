package main

import (
	"GoWebScaffold/infras/starter"
	"GoWebScaffold/infras/utils"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Application Starting  ......")

	var spts []*starter.StarterOption

	// 添加一个异步日志记录器demo
	file, err := os.OpenFile("./log/info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	utils.ErrorHandler(err)
	spts = append(spts, starter.NewStarterOption(starter.LoggerStarterKey, file))

	//2.创建应用程序启动管理器
	application := starter.NewApplication(spts...)
	application.Start()

	fmt.Println(application.Sctx.GetConfig().MysqlConf)

	fmt.Println("Application Running  ......")
}
