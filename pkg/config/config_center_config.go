package config

// ConfigCenterConfig 配置中心配置
type ConfigCenterConfig struct {
	Address       string `yaml:"address" json:"address"`
	Username      string `yaml:"username" json:"username"`
	Password      string `yaml:"password" json:"password"`
	Namespace     string `yaml:"namespace" json:"namespace"`
	Group         string `yaml:"group" json:"group"`
	DataID        string `yaml:"data_id" json:"data_id"`                 // 框架的配置
	ServiceDataID string `yaml:"service_data_id" json:"service_data_id"` // 当前服务的业务配置
}

const CONFIG_CENTER_CONFIG_ENV_KEY = "LM_CONFIG_CENTER_CONFIG"
