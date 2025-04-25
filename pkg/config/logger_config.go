package config

import (
	userLogger "user/pkg/logger"

	"github.com/dubbogo/gost/log/logger"
)

type LoggerConfig struct {
	Level string `mapstructure:"level" yaml:"level" json:"level"`
}

func GetLogger(config *LoggerConfig) (func() logger.Logger, error) {
	Log := func() logger.Logger {
		return userLogger.GetLogger()
	}

	// logger 实际已经使用自定义的替换了框架的，这边使用自定义的
	userLogger.SetLoggerLevel(config.Level)

	return Log, nil
}
