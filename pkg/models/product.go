package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Slug        string `json:"slug" gorm:"unique"`
	Stock       int64  `json:"stock"`
	Price       int64  `json:"price"`
	Description string `json:"description"`
	// Time
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
