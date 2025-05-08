package types

import (
	"sync"
	"time"
)

// Cache 定义缓存接口
type Cache interface {
	BuildKey(key string) string                                                                                  // 构建缓存键
	Exists(key string) (bool, error)                                                                             // 检查缓存是否存在
	Delete(key string) error                                                                                     // 删除缓存
	Flush() error                                                                                                // 删除所有缓存
	GetTTL(key string) (time.Duration, error)                                                                    // 获取缓存剩余时间
	Get(key string, target *string) (bool, error)                                                                // 获取缓存（如果存在则返回 true，否则返回 false）
	Set(key string, value *string, duration time.Duration) error                                                 // 设置缓存（如果 duration 为 0，则使用默认缓存时间）
	Add(key string, value *string, duration time.Duration) error                                                 // 添加缓存（和 Set 一样），如果存在则不处理
	GetOrSet(key string, target *string, callback func() (*string, error), duration time.Duration) (bool, error) // 获取缓存（和 Get 一样），不存在则使用回调函数设置（回调函数返回 nil 时，不会设置缓存）
	GetJson(key string, target any) (bool, error)                                                                // 获取 json 并反序列化（target 为指针，和 json.Unmarshal 一样，如果存在则返回 true，否则返回 false）
	SetJson(key string, value any, duration time.Duration) error                                                 // 设置 json（value 为 json 对象的值或指针，和 json.Marshal 一样）
	AddJson(key string, value any, duration time.Duration) error                                                 // 添加 json（和 SetJson 一样）
	GetOrSetJson(key string, target any, callback func() (any, error), duration time.Duration) (bool, error)     // 获取缓存 json（target 为指针，和 json.Unmarshal 一样）
}

// BaseCache 定义缓存基础结构
type BaseCache struct {
	KeyPrefix       string        // 自定义 key 前缀
	DurationDefault time.Duration // 默认缓存时间
	Mutex           *sync.RWMutex // 用于防止并发
}

// BuildKey 构建缓存键
func (c *BaseCache) BuildKey(key string) string {
	if c.KeyPrefix == "" {
		return key
	}
	return c.KeyPrefix + ":" + key
}
