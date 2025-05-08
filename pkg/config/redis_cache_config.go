package config

import "time"

type CacheRedisConfig struct {
	KeyPrefix     string        `mapstructure:"key-prefix" json:"key-prefix" yaml:"key-prefix"`             // 缓存键前缀
	DurationHuman string        `mapstructure:"duration-human" json:"duration-human" yaml:"duration-human"` // 语义化缓存过期时间，格式：1h，有效单位："ms", "s", "m", "h", "d"
	Duration      time.Duration `mapstructure:"duration" json:"duration" yaml:"duration"`
}
