package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Stock       int64          `json:"stock"`
	Price       int64          `json:"price"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime:true"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}
