package config

import (
	"log"
	"mini-library-backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("✅ Users already seeded.")
		return
	}

	users := []models.User{
		{
			Username: "admin",
			Password: hash("admin123"),
			Role:     "admin",
		},
		{
			Username: "user",
			Password: hash("user123"),
			Role:     "user",
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			log.Printf("❌ Failed to seed user %s: %v", user.Username, err)
		}
	}

	log.Println("✅ Seeded admin and user accounts.")
}

func hash(pw string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Error hashing password: %v", err)
	}
	return string(hashed)
}
