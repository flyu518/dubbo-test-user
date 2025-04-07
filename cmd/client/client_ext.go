package main

import (
	"context"

	"github.com/dubbogo/gost/log/logger"

	"user/api"
)

// Register 用户注册
func (c *Client) Register(ctx context.Context, username, password string) (bool, error) {
	svc, err := c.GetUserService()
	if err != nil {
		return false, err
	}

	req := &api.RegisterRequest{
		Username: username,
		Password: password,
	}

	resp, err := svc.Register(ctx, req)
	if err != nil {
		logger.Errorf("注册用户失败: %v", err)
		return false, err
	}

	return resp.Success, nil
}

// Login 用户登录
func (c *Client) Login(ctx context.Context, username, password string) (*api.User, string, error) {
	svc, err := c.GetUserService()
	if err != nil {
		return nil, "", err
	}

	req := &api.LoginRequest{
		Username: username,
		Password: password,
	}

	resp, err := svc.Login(ctx, req)
	if err != nil {
		logger.Errorf("用户登录失败: %v", err)
		return nil, "", err
	}

	return resp.User, resp.Token, nil
}

// GetUser 获取用户信息
func (c *Client) GetUser(ctx context.Context, username string) (*api.User, error) {
	svc, err := c.GetUserService()
	if err != nil {
		return nil, err
	}

	req := &api.GetUserRequest{
		Username: username,
	}

	resp, err := svc.GetUser(ctx, req)
	if err != nil {
		logger.Errorf("获取用户信息失败: %v", err)
		return nil, err
	}

	return resp.User, nil
}
