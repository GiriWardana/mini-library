package config

import (
	"mini-library-backend/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Borrow{})
}
