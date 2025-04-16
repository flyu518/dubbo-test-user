package main

import (
	"context"
	"os"
	"testing"
	"user/internal/model"
	"user/pkg/global"
	"user/pkg/util"

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
	// 初始化全局变量
	global.InitGlobal("./config/config.yaml")

	// 获取 dubbo 实例和服务端
	instance = util.GetDubboInstance(global.ConfigCenterConfig)
	cli = util.GetDubboClient(instance)

	// 获取服务
	var err error
	srv, err = api.NewUserService(cli)
	if err != nil {
		panic(err)
	}

	// 调用 m.Run 执行测试
	code := m.Run()

	// 删除测试用户
	global.DB.Where("username = ?", "单元测试生成").Delete(&model.User{})

	os.Exit(code)
}

func TestUser(t *testing.T) {
	t.Run("注册", func(t *testing.T) {
		res, err := srv.Register(context.Background(), &api.RegisterRequest{
			Username: "单元测试生成",
			Password: "123456",
		})

		assert.NoError(t, err)

		assert.Equal(t, true, res.Success)
	})

	t.Run("获取用户", func(t *testing.T) {
		res, err := srv.GetUser(context.Background(), &api.GetUserRequest{
			Username: "单元测试生成",
		})

		assert.NoError(t, err)

		assert.Equal(t, "单元测试生成", res.User.Username)
	})
}
