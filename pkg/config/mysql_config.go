package config

import "fmt"

// MysqlConfig Mysql配置结构
type MysqlConfig struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn 构建MySQL连接字符串
func (c *MysqlConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DbName,
		c.Config,
	)
}
