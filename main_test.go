package main

import (
	"context"
	"os"
	"testing"
	"user/pkg"

	"dubbo.apache.org/dubbo-go/v3"
	"dubbo.apache.org/dubbo-go/v3/client"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/flyu518/dubbo-test-sdk/user/api"
	"github.com/stretchr/testify/assert"
)

var instance *dubbo.Instance
var cli *client.Client
var srv api.UserService

func TestMain(m *testing.M) {
	// 获取 dubbo 实例和服务端
	instance = pkg.GetDubboInstance(pkg.ParseEnvCenterConfig())
	cli = pkg.GetClient(instance)

	// 获取服务
	var err error
	srv, err = api.NewUserService(cli)
	if err != nil {
		panic(err)
	}

	// 调用 m.Run 执行测试
	code := m.Run()

	os.Exit(code)
}

func TestUser(t *testing.T) {
	t.Run("获取用户", func(t *testing.T) {
		res, err := srv.GetUser(context.Background(), &api.GetUserRequest{
			Username: "123456",
		})

		assert.NoError(t, err)

		assert.Equal(t, "123456", res.User.Username)
	})
}
