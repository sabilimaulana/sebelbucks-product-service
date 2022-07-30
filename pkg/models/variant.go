package models

import (
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Stock       int64          `json:"stock"`
	Price       int64          `json:"price"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
