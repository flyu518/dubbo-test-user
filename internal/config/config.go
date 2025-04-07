package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 全局单例实例
var (
	globalConfig *Config
	globalViper  *viper.Viper
	viperOnce    sync.Once
	mysqlDB      *gorm.DB
	redisClient  redis.UniversalClient
	initialized  bool
)

// Config 总配置结构
type Config struct {
	MySQL MySQLConfig `mapstructure:"mysql" yaml:"mysql" json:"mysql"`
	Redis RedisConfig `mapstructure:"redis" yaml:"redis" json:"redis"`
}

// InitConfig 初始化配置
func InitConfig(configPath string) *Config {
	viperOnce.Do(func() {
		v := viper.New()
		v.SetConfigFile(configPath)
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("读取配置文件失败: %s", err))
		}

		// 监控配置文件变化
		// v.WatchConfig()
		// v.OnConfigChange(func(e fsnotify.Event) {
		// 	fmt.Printf("配置文件变更: %s\n", e.Name)
		// 	if err := v.Unmarshal(&globalConfig); err != nil {
		// 		fmt.Printf("配置重新加载失败: %s\n", err)
		// 	}
		// })

		// 解析配置
		globalConfig = &Config{}
		if err := v.Unmarshal(globalConfig); err != nil {
			panic(fmt.Errorf("解析配置文件失败: %s", err))
		}

		globalViper = v
	})

	return globalConfig
}

// GetConfig 获取全局配置
func GetConfig() *Config {
	if globalConfig == nil {
		panic("配置未初始化，请先调用 InitConfig")
	}
	return globalConfig
}

// GetViper 获取全局Viper实例
func GetViper() *viper.Viper {
	if globalViper == nil {
		panic("配置未初始化，请先调用 InitConfig")
	}
	return globalViper
}

// GetMySQL 获取全局MySQL连接
func GetMySQL() *gorm.DB {
	if !initialized {
		panic("请先调用 Setup 初始化配置")
	}
	return mysqlDB
}

// GetRedis 获取全局Redis连接
func GetRedis() redis.UniversalClient {
	if !initialized {
		panic("请先调用 Setup 初始化配置")
	}
	return redisClient
}

// Setup 完整的应用程序配置初始化
func Setup(configPath string) error {
	// 1. 加载配置文件
	config := InitConfig(configPath)

	// 2. 初始化MySQL
	var err error
	mysqlDB, err = config.MySQL.GetDB()
	if err != nil {
		return fmt.Errorf("初始化MySQL失败: %v", err)
	}

	// 3. 初始化Redis
	redisClient, err = config.Redis.GetClient()
	if err != nil {
		// 关闭之前初始化的连接
		if mysqlDB != nil {
			config.MySQL.Close()
		}
		return fmt.Errorf("初始化Redis失败: %v", err)
	}

	initialized = true
	return nil
}

// Close 关闭所有连接
func Close() {
	if initialized {
		cfg := GetConfig()

		// 关闭MySQL连接
		if err := cfg.MySQL.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "关闭MySQL连接失败: %v\n", err)
		}

		// 关闭Redis连接
		if err := cfg.Redis.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "关闭Redis连接失败: %v\n", err)
		}

		mysqlDB = nil
		redisClient = nil
		initialized = false
	}
}
