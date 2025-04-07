package main

import (
	"errors"
	"os"
	"sync"

	"dubbo.apache.org/dubbo-go/v3"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"user/api"
)

// 全局变量
var (
	dubboClient     *dubbo.Instance
	userServiceOnce sync.Once
	userService     api.UserService
	initialized     bool
)

// Client 是用户服务的客户端管理器
type Client struct {
	configPath string
}

// NewClient 创建一个新的客户端管理器
func NewClient(configPath string) *Client {
	return &Client{
		configPath: configPath,
	}
}

// Init 初始化客户端
func (c *Client) Init() error {
	if initialized {
		return nil
	}

	// 设置配置文件路径
	if c.configPath != "" {
		os.Setenv("DUBBO_GO_CONFIG_PATH", c.configPath)
	}

	// 初始化dubbo
	var err error
	dubboClient, err = dubbo.NewInstance()
	if err != nil {
		return err
	}
	if err = dubboClient.Start(); err != nil {
		return err
	}

	// 标记为已初始化
	initialized = true
	return nil
}

// GetUserService 获取用户服务客户端
func (c *Client) GetUserService() (api.UserService, error) {
	if !initialized {
		return nil, errors.New("客户端未初始化，请先调用Init方法")
	}

	var err error
	userServiceOnce.Do(func() {
		// 引用远程服务
		userService, err = api.NewUserService(dubbo.NewConsumerClient())
	})

	if err != nil {
		return nil, err
	}

	return userService, nil
}

// Close 关闭客户端连接
func (c *Client) Close() error {
	if initialized && dubboClient != nil {
		dubboClient.Stop()
		initialized = false
	}
	return nil
}
