package main

import (
	"context"

	"github.com/dubbogo/gost/log/logger"
)

func main() {
	// 创建客户端
	client := NewClient("./config/dubbogo.yaml")

	// 初始化
	if err := client.Init(); err != nil {
		logger.Errorf("初始化客户端失败: %v", err)
		return
	}

	// 确保资源释放
	defer client.Close()

	// 示例：获取用户信息
	ctx := context.Background()
	user, err := client.GetUser(ctx, "test_user")
	if err != nil {
		logger.Errorf("获取用户失败: %v", err)
		return
	}

	logger.Infof("用户信息: %v", user)
}
