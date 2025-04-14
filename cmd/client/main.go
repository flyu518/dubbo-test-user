package main

import (
	"context"

	"user/pkg"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/user/api"
)

func main() {
	// 通过环境变量获取配置
	nacosConfig := pkg.GetNacosConfig()
	centerConfig := pkg.GetCenterConfig()

	// 获取 dubbo 实例和客户端
	instance := pkg.GetDubboInstance(pkg.GetConfigCenterOption(nacosConfig, centerConfig))
	client := pkg.GetClient(instance)

	// 获取服务
	srv, err := api.NewUserService(client)
	if err != nil {
		panic(err)
	}

	logger.Infof("用户客户端已启动")

	res, err := srv.Register(context.Background(), &api.RegisterRequest{
		Username: "test",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}

	logger.Infof("注册结果: %v", res)
}
