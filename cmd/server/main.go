package main

import (
	"os"
	"os/signal"
	"syscall"

	"user/internal/config"
	"user/internal/handler"
	"user/pkg"

	"github.com/flyu518/dubbo-test-sdk/user/api"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
)

// 启动应用
func main() {
	root_path := os.Getenv("DUBBO_GO_ROOT_PATH")
	config.InitConfig(root_path + "/config/config.yaml") // 地址待处理

	// 配置中心启动方式 ------------
	nacosConfig := pkg.NacosConfig{
		Address:  "127.0.0.1:8848",
		Username: "",
		Password: "",
	}
	centerConfig := pkg.CenterConfig{
		Namespace: "dev",
		Group:     "user",
		DataID:    "server.yaml",
	}

	// 获取 dubbo 实例和服务端
	instance := pkg.GetDubboInstance(pkg.GetConfigCenterOption(&nacosConfig, &centerConfig))
	srv := pkg.GetServer(instance)

	// 注册服务
	if err := api.RegisterUserServiceHandler(srv, &handler.UserHandler{}); err != nil {
		panic(err)
	}

	logger.Infof("用户服务已启动")

	// 启动服务
	if err := srv.Serve(); err != nil {
		logger.Error(err)
	}

	// 读取本地配置启动方式 ------------
	// api.SetProviderUserService(new(service.UserService))
	// if err := dubbo.Load(); err != nil {
	// 	panic(err)
	// }

	// logger.Infof("用户服务已启动")

	// waitForShutdown()
}

// 优雅退出处理
func waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	logger.Infof("收到信号: %s, 退出应用", sig)
}
