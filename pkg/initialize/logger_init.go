package initialize

import (
	"user/pkg/config"
	pkgLogger "user/pkg/logger"

	"github.com/dubbogo/gost/log/logger"
)

// InitLogger 初始化日志
func InitLogger(config *config.Config) (func() logger.Logger, error) {
	Log := func() logger.Logger {
		return pkgLogger.GetLogger()
	}

	// // 如果自定义配置比较多，并且没有单独的设置方法，可以先获取默认配置，然后修改后重新初始化
	// // 获取默认配置
	// encoderConfig := pkgLogger.GetZapEncoderConfigDefault()
	// zapConfig := pkgLogger.GetZapConfigDefault(encoderConfig)

	// // 初始化日志
	// pkgLogger.InitLogger(&pkgLogger.Config{
	// 	ZapConfig: zapConfig,
	// })

	// 设置日志等级
	pkgLogger.SetLoggerLevel(config.Logger.Level)

	// 注入服务名
	pkgLogger.InjectServiceName(config.System.ServiceName)

	return Log, nil
}
