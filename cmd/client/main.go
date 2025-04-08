package main

import (
	"context"

	"dubbo.apache.org/dubbo-go/v3"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/flyu518/dubbo-test-sdk/user/api"
)

var srv = new(api.UserServiceImpl)

func main() {
	api.SetConsumerUserService(srv)
	if err := dubbo.Load(); err != nil {
		panic(err)
	}

	logger.Infof("用户服务已启动")

	res, err := srv.Register(context.Background(), &api.RegisterRequest{
		Username: "test",
		Password: "123456",
	})
	if err != nil {
		panic(err)
	}

	logger.Infof("注册结果: %v", res)
}
