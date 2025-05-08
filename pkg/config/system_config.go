package config

type SystemConfig struct {
	ServiceName string `mapstructure:"service-name" yaml:"service-name" json:"service-name"`
	Env         Env    `mapstructure:"env" yaml:"env" json:"env"`
}

type Env string

const (
	EnvLocal Env = "local"
	EnvDev   Env = "dev"
	EnvProd  Env = "prod"
)
