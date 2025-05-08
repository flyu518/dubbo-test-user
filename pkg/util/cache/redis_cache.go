package cache

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"user/pkg/types"

	"github.com/redis/go-redis/v9"
)

// CacheRedis 定义Redis缓存
type CacheRedis struct {
	types.BaseCache
	Redis redis.UniversalClient
}

// NewCacheRedis 创建 Redis 缓存
func NewCacheRedis(keyPrefix string, durationDefault time.Duration, redis redis.UniversalClient) *CacheRedis {
	return &CacheRedis{
		BaseCache: types.BaseCache{
			KeyPrefix:       keyPrefix,
			DurationDefault: durationDefault,
			Mutex:           &sync.RWMutex{},
		},
		Redis: redis,
	}
}

// Exists 检查缓存是否存在
func (c *CacheRedis) Exists(key string) (bool, error) {
	exists, err := c.Redis.Exists(context.Background(), c.BuildKey(key)).Result()
	return exists > 0, err
}

// Delete 删除缓存
func (c *CacheRedis) Delete(key string) error {
	return c.Redis.Del(context.Background(), c.BuildKey(key)).Err()
}

// Flush 删除所有缓存
func (c *CacheRedis) Flush() error {
	return c.Redis.FlushAll(context.Background()).Err()
}

// GetTTL 获取缓存过期时间（单位秒，-1表示永久，-2表示不存在）
func (c *CacheRedis) GetTTL(key string) (time.Duration, error) {
	return c.Redis.TTL(context.Background(), c.BuildKey(key)).Result()
}

// Get 获取缓存（如果存在则返回 true，否则返回 false）
func (c *CacheRedis) Get(key string, target *string) (bool, error) {
	value, err := c.Redis.Get(context.Background(), c.BuildKey(key)).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	*target = value

	return true, nil
}

// Set 设置缓存（如果 duration 为 0，则使用默认缓存时间）
func (c *CacheRedis) Set(key string, value *string, duration time.Duration) error {
	if duration == 0 {
		duration = c.DurationDefault
	}
	return c.Redis.Set(context.Background(), c.BuildKey(key), value, duration).Err()
}

// Add 添加缓存（和 Set 一样），如果存在则不处理
func (c *CacheRedis) Add(key string, value *string, duration time.Duration) error {
	return c.Redis.SetNX(context.Background(), c.BuildKey(key), value, duration).Err()
}

// GetOrSet 获取缓存（和 Get 一样），不存在则使用回调函数设置
func (c *CacheRedis) GetOrSet(key string, target *string, callback func() (*string, error), duration time.Duration) (bool, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 尝试从缓存获取
	var value string
	if found, err := c.Get(key, &value); err != nil {
		return false, err
	} else if found {
		*target = value
		return true, nil
	}

	// 缓存中没有值，调用回调函数获取数据
	val, err := callback()
	if err != nil {
		return false, err
	}

	if val == nil {
		return false, nil
	}

	// 设置缓存
	if err := c.Set(key, val, duration); err != nil {
		return false, err
	}

	*target = *val
	return true, nil
}

// GetJson 获取 json 并反序列化（target 为指针，和 json.Unmarshal 一样，如果存在则返回 true，否则返回 false）
func (c *CacheRedis) GetJson(key string, target any) (bool, error) {
	var value string
	if found, err := c.Get(key, &value); err != nil {
		return false, err
	} else if !found {
		return false, nil
	}

	return true, json.Unmarshal([]byte(value), target)
}

// SetJson 设置 json（value 为 json 对象的值或指针，和 json.Marshal 一样）
func (c *CacheRedis) SetJson(key string, value any, duration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	jsonStr := string(json)
	return c.Set(key, &jsonStr, duration)
}

// AddJson 添加 json（和 SetJson 一样），如果存在则不处理
func (c *CacheRedis) AddJson(key string, value any, duration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}
	jsonStr := string(json)
	return c.Add(key, &jsonStr, duration)
}

// GetOrSetJson 获取缓存 json（target 为指针，和 json.Unmarshal 一样）
func (c *CacheRedis) GetOrSetJson(key string, target any, callback func() (any, error), duration time.Duration) (bool, error) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	// 尝试从缓存获取
	if found, err := c.GetJson(key, target); err != nil {
		return false, err
	} else if found {
		return true, nil
	}

	// 缓存中没有值，调用回调函数获取数据
	val, err := callback()
	if err != nil {
		return false, err
	}
	if val == nil {
		return false, nil
	}

	// 设置缓存
	if err := c.SetJson(key, val, duration); err != nil {
		return false, err
	}

	// 将val的数据转换为JSON再解析到target
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return false, err
	}
	return true, json.Unmarshal(jsonBytes, target)
}
