/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package service

import (
	"context"

	"github.com/flyu518/dubbo-test-sdk/user/api"

	"github.com/dubbogo/gost/log/logger"
)

// UserService 实现用户服务
type UserService struct {
}

// Register 实现用户注册服务
func (u *UserService) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {

	logger.Infof("收到注册请求: %v", req.Username)

	// 实际应用中，这里应该有真正的业务逻辑实现
	// 这里简单返回成功
	return &api.RegisterResponse{
		Success: true,
	}, nil
}

// Login 实现用户登录服务
func (u *UserService) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {

	logger.Infof("收到登录请求: %v", req.Username)

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
func (u *UserService) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {

	logger.Infof("收到获取用户信息请求: %v", req.Username)

	// 实际应用中，这里应该根据用户名查询数据库
	// 这里简单返回一个模拟用户
	return &api.GetUserResponse{
		User: &api.User{
			Username: req.Username,
			Email:    req.Username + "@example.com",
			Phone:    "12345678901",
		},
	}, nil
}
