package db

import (
	"log"

	"github.com/sabilimaulana/sebelbucks-product-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	// Make connection
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	// Error handling
	// Will stop the program
	if err != nil {
		log.Fatalln(err)
	}

	// Migrate tables
	db.AutoMigrate(&models.Product{}, &models.Variant{})

	return Handler{db}
}
