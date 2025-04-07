package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// MySQLConfig MySQL配置结构
type MySQLConfig struct {
	Host         string `mapstructure:"host" yaml:"host" json:"host"`
	Port         string `mapstructure:"port" yaml:"port" json:"port"`
	DBName       string `mapstructure:"db-name" yaml:"db-name" json:"db-name"`
	Username     string `mapstructure:"username" yaml:"username" json:"username"`
	Password     string `mapstructure:"password" yaml:"password" json:"password"`
	Singular     bool   `mapstructure:"singular" yaml:"singular" json:"singular"`
	Config       string `mapstructure:"config" yaml:"config" json:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" yaml:"max-idle-conns" json:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" yaml:"max-open-conns" json:"max-open-conns"`
	LogMode      string `mapstructure:"log-mode" yaml:"log-mode" json:"log-mode"`
	LogZap       bool   `mapstructure:"log-zap" yaml:"log-zap" json:"log-zap"`

	db *gorm.DB // 缓存的数据库连接
}

// DSN 构建MySQL连接字符串
func (c *MySQLConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.Config,
	)
}

// GetDB 获取MySQL连接（懒加载）
func (c *MySQLConfig) GetDB() (*gorm.DB, error) {
	if c.db != nil {
		return c.db, nil
	}

	// 设置日志级别
	var logLevel logger.LogLevel
	switch c.LogMode {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Error
	}

	// 设置GORM配置
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: c.Singular,
		},
	}

	// 创建连接
	db, err := gorm.Open(mysql.Open(c.DSN()), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("连接MySQL失败: %s", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	c.db = db
	return db, nil
}

// Close 关闭数据库连接
func (c *MySQLConfig) Close() error {
	if c.db != nil {
		sqlDB, err := c.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
