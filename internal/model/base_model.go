package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"column:id;type:int(20);primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;index"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
