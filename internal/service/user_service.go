package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"user/internal/model"
	"user/pkg/global"

	"github.com/flyu518/dubbo-test-sdk/user/api"
)

var UserService *userService

// userService 实现用户服务
type userService struct {
}

// Register 实现用户注册服务
func (u *userService) Register(req *api.RegisterRequest) (*api.RegisterResponse, error) {
	user := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	// 检查用户是否存在
	var count int64
	if err := global.DB.Model(&model.User{}).Where("username = ?", user.Username).Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, errors.New("用户已存在")
	}

	if err := global.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &api.RegisterResponse{
		Success: true,
	}, nil
}

// Login 实现用户登录服务
func (u *userService) Login(req *api.LoginRequest) (*api.LoginResponse, error) {
	// 实际应用中，这里应该验证用户名和密码，查询数据库等
	// 这里简单返回一个模拟用户
	return &api.LoginResponse{
		User: &api.User{
			Username: req.Username,
			Email:    req.Username + "@example.com",
			Phone:    "12345678901",
		},
		Token: "模拟的JWT令牌",
	}, nil
}

// GetUser 实现获取用户信息服务
func (u *userService) GetUser(req *api.GetUserRequest) (*api.GetUserResponse, error) {
	// 仅仅为了做 redis 测试
	global.Log.Info("开始缓存", req.Username)
	global.Redis.Set(context.Background(), req.Username, "123456", 10*time.Second)
	username, err := global.Redis.Get(context.Background(), req.Username).Result()
	if err != nil {
		global.Log.Error("获取缓存失败", err)
		return nil, err
	}
	global.Log.Info("获取缓存", username)

	// 读取数据库
	global.Log.Info("开始读取数据库", req.Username)
	user := model.User{}

	if err := global.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		global.Log.Error("读取数据库失败", err)
		return nil, err
	}

	userJson, _ := json.Marshal(user)
	global.Log.Info("读取数据库成功", string(userJson))

	return &api.GetUserResponse{
		User: &api.User{
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
		},
	}, nil
}
