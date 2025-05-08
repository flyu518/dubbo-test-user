package main

import (
	"user/internal/handler"
	internalInitialize "user/internal/initialize"
	"user/pkg/global"
	"user/pkg/initialize"
	"user/pkg/util"

	"dubbo.apache.org/dubbo-go/v3/server"
	"github.com/flyu518/dubbo-test-sdk/user/api"
)

// 启动应用
func main() {
	// 初始化全局变量
	initialize.Init()

	// 初始化私有数据
	internalInitialize.Init()

	// 获取 dubbo 实例和服务端
	instance := util.GetDubboInstance(global.ConfigCenterConfig)
	srv := util.GetDubboServer(instance)

	// 注册服务，这种方式启动的，filter 只能这样写，不能在 yaml 中配置
	err := api.RegisterUserServiceHandler(srv, handler.GetUserHandler(instance), server.WithFilter("otelServerTrace,logTraceFilter"))
	if err != nil {
		panic(err)
	}

	global.Log.Info("用户服务已启动")

	// 启动服务
	if err := srv.Serve(); err != nil {
		global.Log.Error(err)
	}
}
