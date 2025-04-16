package util

import (
	"user/internal/model"

	"gorm.io/gorm"
)

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
