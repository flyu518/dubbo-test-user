package handler

import (
	"context"
	"user/internal/service"
	"user/pkg/global"

	"github.com/flyu518/dubbo-test-sdk/user/api"

	"github.com/dubbogo/gost/log/logger"
)

func GetUserHandler() *UserHandler {
	return &UserHandler{}
}

// UserHandler 实现用户服务
type UserHandler struct {
}

// Register 实现用户注册服务
func (u *UserHandler) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	//
	logger.Infof("收到注册请求: %v", req.Username)
	global.Log().Infof("收到注册请求222: %v", req.Username)

	return service.UserService.Register(req)
}

// Login 实现用户登录服务
func (u *UserHandler) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {

	logger.Infof("收到登录请求: %v", req.Username)

	return service.UserService.Login(req)
}

// GetUser 实现获取用户信息服务
func (u *UserHandler) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {

	logger.Infof("收到获取用户信息请求: %v", req.Username)

	return service.UserService.GetUser(req)
}
