package main

import (
	"context"

	"user/pkg"

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/user/api"
)

func main() {
	// 配置中心启动方式 ------------
	nacosConfig := pkg.NacosConfig{
		Address:  "127.0.0.1:8848",
		Username: "",
		Password: "",
	}
	centerConfig := pkg.CenterConfig{
		Namespace: "dev",
		Group:     "user",
		DataID:    "client.yaml",
	}

	// 获取 dubbo 实例和客户端
	instance := pkg.GetDubboInstance(pkg.GetConfigCenterOption(&nacosConfig, &centerConfig))
	client := pkg.GetClient(instance)

	// 获取服务
	srv, err := api.NewUserService(client)
	if err != nil {
		panic(err)
	}

	// 读取本地配置启动方式 ------------
	// var srv = new(api.UserServiceImpl)
	// api.SetConsumerUserService(srv)
	// if err := dubbo.Load(); err != nil {
	// 	panic(err)
	// }

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
