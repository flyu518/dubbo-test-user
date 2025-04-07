package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConfig Redis配置结构
type RedisConfig struct {
	Mode     string       `mapstructure:"mode" yaml:"mode" json:"mode"`
	Password string       `mapstructure:"password" yaml:"password" json:"password"`
	Single   RedisSingle  `mapstructure:"single" yaml:"single" json:"single"`
	Cluster  RedisCluster `mapstructure:"cluster" yaml:"cluster" json:"cluster"`

	client redis.UniversalClient // 缓存的Redis客户端
}

// RedisSingle 单节点配置
type RedisSingle struct {
	Addr string `mapstructure:"addr" yaml:"addr" json:"addr"`
	DB   int    `mapstructure:"db" yaml:"db" json:"db"`
}

// RedisCluster 集群配置
type RedisCluster struct {
	Addrs []string `mapstructure:"addrs" yaml:"addrs" json:"addrs"`
}

// GetClient 获取Redis客户端（懒加载）
func (c *RedisConfig) GetClient() (redis.UniversalClient, error) {
	if c.client != nil {
		return c.client, nil
	}

	var client redis.UniversalClient

	switch c.Mode {
	case "single":
		// 单节点模式
		client = redis.NewClient(&redis.Options{
			Addr:     c.Single.Addr,
			Password: c.Password,
			DB:       c.Single.DB,
		})
	case "cluster":
		// 集群模式
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    c.Cluster.Addrs,
			Password: c.Password,
		})
	default:
		return nil, fmt.Errorf("不支持的Redis模式: %s", c.Mode)
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("Redis连接失败: %s", err)
	}

	c.client = client
	return client, nil
}

// Close 关闭Redis连接
func (c *RedisConfig) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}
