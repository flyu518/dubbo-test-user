package main

import (
	"user/internal/handler"
	"user/pkg"

	"github.com/flyu518/dubbo-test-sdk/user/api"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
)

// 启动应用
func main() {
	// 获取 dubbo 实例和服务端
	instance := pkg.GetDubboInstance(pkg.ParseEnvCenterConfig())
	srv := pkg.GetServer(instance)

	// 注册服务
	if err := api.RegisterUserServiceHandler(srv, &handler.UserHandler{}); err != nil {
		panic(err)
	}

	logger.Info("用户服务已启动")

	// 启动服务
	if err := srv.Serve(); err != nil {
		logger.Error(err)
	}
}
