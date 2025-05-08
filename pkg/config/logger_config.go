package config

type LoggerConfig struct {
	Level string `mapstructure:"level" yaml:"level" json:"level"`
}
