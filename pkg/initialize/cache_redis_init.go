package initialize

import (
	"fmt"
	"user/pkg/config"
	"user/pkg/util"
	"user/pkg/util/cache"

	"github.com/redis/go-redis/v9"
)

func InitCacheRedis(config *config.Config, redis redis.UniversalClient) (*cache.CacheRedis, error) {
	if redis == nil {
		return nil, fmt.Errorf("redis 未初始化")
	}

	duration, err := util.ParseDuration(config.CacheRedis.DurationHuman)
	if err != nil {
		return nil, fmt.Errorf("解析缓存过期时间失败: %v", err)
	}

	key_prefix := config.CacheRedis.KeyPrefix
	if key_prefix == "" {
		key_prefix = config.System.ServiceName
	}

	return cache.NewCacheRedis(key_prefix, duration, redis), nil
}
