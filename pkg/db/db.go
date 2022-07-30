package db

import (
	"log"

	"github.com/sabilimaulana/sebelbucks-product-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	// Make connection
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	// Error handling
	// Will stop the program
	if err != nil {
		log.Fatalln(err)
	}

	// Migrate tables
	db.AutoMigrate(&models.Product{}, &models.Variant{})

	return Handler{db}
}
