package initialize

import (
	"encoding/json"
	"fmt"
	"os"
	"user/pkg/config"
)

// InitEnvConfigCenterConfig 初始化环境变量里面的配置中心配置
func InitEnvConfigCenterConfig() *config.ConfigCenterConfig {
	center_config := os.Getenv(config.CONFIG_CENTER_CONFIG_ENV_KEY)
	if center_config == "" {
		panic(fmt.Sprintf("环境变量 %s 未设置", config.CONFIG_CENTER_CONFIG_ENV_KEY))
	}

	configCenterConfig := config.ConfigCenterConfig{}
	err := json.Unmarshal([]byte(center_config), &configCenterConfig)
	if err != nil {
		panic(fmt.Sprintf("环境变量 %s 解析失败: %s", config.CONFIG_CENTER_CONFIG_ENV_KEY, err.Error()))
	}

	return &configCenterConfig
}
