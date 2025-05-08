package initialize

import (
	"user/pkg/config"
	"user/pkg/global"
	_ "user/pkg/imports" // 导入自定义依赖
)

func Init() {
	var err error

	global.ConfigCenterConfig = InitEnvConfigCenterConfig()

	global.Config = &config.Config{}
	err = InitConfig(global.ConfigCenterConfig, global.Config)
	if err != nil {
		panic(err)
	}

	global.DB, err = InitMysql(global.Config)
	if err != nil {
		panic(err)
	}

	global.Redis, err = InitRedis(global.Config)
	if err != nil {
		panic(err)
	}

	global.Log, global.LogCtx, err = InitLogger(global.Config)
	if err != nil {
		panic(err)
	}

	global.Cache, err = InitCacheRedis(global.Config, global.Redis)
	if err != nil {
		panic(err)
	}
}
