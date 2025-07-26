package main

import (
	"log"
	"mini-library-backend/controllers"
	"mini-library-backend/config"
	"mini-library-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load() // loads .env into os.Environ
	// Inisialisasi konfigurasi (DB)
	
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	controllers.SetDB(db)
	// ✅ Correct DB injection for each controller
	controllers.SetBookDB(db)
	controllers.SetUserDB(db)
	controllers.SetBorrowDB(db) // ✅ THIS IS WHAT YOU MISSED

	controllers.SeedDefaultUsers()

	// Optional: auto migrate
	config.AutoMigrate(db)

	// Inisialisasi Gin router
	r := gin.Default()

	// Register routes
	routes.SetupRoutes(r)

	// Run server
	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
