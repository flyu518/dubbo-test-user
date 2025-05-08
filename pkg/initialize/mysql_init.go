package initialize

import (
	"errors"
	"user/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitMysql 初始化MySQL连接
func InitMysql(config *config.Config) (*gorm.DB, error) {
	if config.MySQL.DbName == "" {
		return nil, errors.New("数据库名不能为空")
	}

	mysqlConfig := mysql.Config{
		DSN:                       config.MySQL.Dsn(), // DSN data source name
		DefaultStringSize:         191,                // string 类型字段的默认长度
		SkipInitializeWithVersion: false,              // 根据版本自动配置
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(config.MySQL.LogLevel()),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.MySQL.Prefix,
			SingularTable: config.MySQL.Singular,
		},
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		return nil, err
	}

	db.InstanceSet("gorm:table_options", "ENGINE="+config.MySQL.Engine)
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.MySQL.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MySQL.MaxOpenConns)
	return db, nil
}
