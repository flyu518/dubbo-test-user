package initialize

import (
	"user/internal/model"
	"user/pkg/config"
	"user/pkg/global"

	"gorm.io/gorm"
)

func InitMigrate() {
	// 临时操作（正式环境不要这样）
	if global.Config.System.Env != config.EnvProd {
		if err := MigrateTables(global.DB); err != nil {
			panic(err)
		}
	}
}

// MigrateTables 迁移表
func MigrateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.User{},
	)

	// TODO:: 后续添加上表注释

	if err != nil {
		return err
	}

	return nil
}
