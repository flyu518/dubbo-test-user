package initialize

import (
	"context"
	"fmt"
	"time"

	"user/pkg/config"

	"github.com/redis/go-redis/v9"
)

// InitRedis 初始化Redis客户端
func InitRedis(config *config.Config) (redis.UniversalClient, error) {
	var client redis.UniversalClient

	switch config.Redis.Mode {
	case "single":
		// 单节点模式
		client = redis.NewClient(&redis.Options{
			Addr:     config.Redis.Single.Addr,
			Password: config.Redis.Password,
			DB:       config.Redis.Single.DB,
		})
	case "cluster":
		// 集群模式
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    config.Redis.Cluster.Addrs,
			Password: config.Redis.Password,
		})
	default:
		return nil, fmt.Errorf("不支持的Redis模式: %s", config.Redis.Mode)
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("Redis连接失败: %s", err)
	}

	return client, nil
}
