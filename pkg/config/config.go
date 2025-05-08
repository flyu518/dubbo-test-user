package config

type Config struct {
	System     SystemConfig     `mapstructure:"system" yaml:"system" json:"system"`
	MySQL      MysqlConfig      `mapstructure:"mysql" yaml:"mysql" json:"mysql"`
	Redis      RedisConfig      `mapstructure:"redis" yaml:"redis" json:"redis"`
	Logger     LoggerConfig     `mapstructure:"logger" yaml:"logger" json:"logger"`
	CacheRedis CacheRedisConfig `mapstructure:"cache-redis" yaml:"cache-redis" json:"cache-redis"`
}
