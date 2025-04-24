package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"user/internal/handler"
	"user/internal/initialize"
	"user/pkg/global"

	"dubbo.apache.org/dubbo-go/v3"
	"github.com/flyu518/dubbo-test-sdk/user/api"
)

// 启动应用
func main() {
	// 初始化全局变量
	global.InitGlobal("./config/config.yaml") // 考虑设置个绝对地址
	initialize.Init()

	// // // 获取 dubbo 实例和服务端
	// instance := util.GetDubboInstance(global.ConfigCenterConfig)
	// srv := util.GetDubboServer(instance)

	// // 注册服务
	// if err := api.RegisterUserServiceHandler(srv, handler.GetUserHandler(instance)); err != nil {
	// 	panic(err)
	// }

	// global.Log = logger.GetLogger() // 实例化之后设置，不要在实例化之前设置

	// global.Log.Info("用户服务已启动")

	// // 启动服务
	// if err := srv.Serve(); err != nil {
	// 	global.Log.Error(err)
	// }

	api.SetProviderUserService(handler.GetUserHandler())
	if err := dubbo.Load(dubbo.WithPath("./config/server.yaml")); err != nil {
		panic(err)
	}

	//global.Log = logger.GetLogger() // 实例化之后设置，不要在实例化之前设置
	global.Log().Info("用户服务已启动")

	waitForShutdown()
}

// 优雅退出处理
func waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	fmt.Printf("收到信号: %s, 退出应用", sig)
}
