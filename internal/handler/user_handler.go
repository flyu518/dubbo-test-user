package handler

import (
	"context"
	"time"
	"user/internal/service"
	"user/pkg/global"

	"dubbo.apache.org/dubbo-go/v3"
	"github.com/flyu518/dubbo-test-sdk/user/api"
)

func GetUserHandler(instance *dubbo.Instance) *UserHandler {
	return &UserHandler{}
}

// UserHandler 实现用户服务
type UserHandler struct{}

// Register 实现用户注册服务
func (u *UserHandler) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	// 日志使用方式示例
	global.Log().Infof("收到注册请求: %v", req.Username)

	// 缓存使用方式示例
	var username string
	found, err := global.Cache.GetOrSet("username", &username, func() (*string, error) {
		return &req.Username, nil
	}, 100*time.Second)
	if err != nil {
		global.Log().Error("获取缓存失败", err)
	} else if found {
		global.Log().Infof("获取缓存成功: %v", username)
	} else {
		global.Log().Infof("获取缓存为空: %v", username)
	}

	// 缓存使用方式示例 - json
	global.Cache.SetJson("username-json", req, 100*time.Second)
	req2 := &api.RegisterRequest{}
	found, err = global.Cache.GetOrSetJson("username-json5", req2, func() (any, error) {
		return req, nil
	}, 100*time.Second)
	if err != nil {
		global.Log().Error("获取缓存失败-json：%v", err)
	} else if found {
		global.Log().Infof("获取缓存成功-json： %v", req2)
	} else {
		global.Log().Infof("获取缓存为空-json： %v", req2)
	}

	return service.UserService.Register(req)
}

// Login 实现用户登录服务
func (u *UserHandler) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {

	global.Log().Infof("收到登录请求: %v", req.Username)

	return service.UserService.Login(req)
}

// GetUser 实现获取用户信息服务
func (u *UserHandler) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {

	global.Log().Infof("收到获取用户信息请求: %v", req.Username)

	return service.UserService.GetUser(req)
}
