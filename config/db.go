package config

import (
	"FeedbackAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=myuser password=mypassword dbname=feedbackapidb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func MigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&models.Customer{}, &models.Feedback{})
}
