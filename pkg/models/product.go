package models

import "time"

type Product struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Stock       int64     `json:"stock"`
	Price       int64     `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime:true"`
}
