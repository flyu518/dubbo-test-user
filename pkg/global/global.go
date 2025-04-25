package global

import (
	"user/pkg/config"
	_ "user/pkg/imports" // 导入自定义依赖

	_ "dubbo.apache.org/dubbo-go/v3/imports" // 导入dubbo-go的依赖，必须的
	"github.com/dubbogo/gost/log/logger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ConfigCenterConfig *config.ConfigCenterConfig

	Config *config.Config
	DB     *gorm.DB
	Redis  redis.UniversalClient
	Log    func() logger.Logger
)

func InitGlobal(configPath string) {
	var err error

	ConfigCenterConfig = config.GetEnvConfigCenterConfig()

	Config = &config.Config{}
	err = config.GetConfig(configPath, Config)
	if err != nil {
		panic(err)
	}

	DB, err = config.GetMysql(&Config.MySQL)
	if err != nil {
		panic(err)
	}

	Redis, err = config.GetRedis(&Config.Redis)
	if err != nil {
		panic(err)
	}

	Log, err = config.GetLogger(&Config.Logger)
	if err != nil {
		panic(err)
	}
}
