package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/kayki-araujo/lynk-api/internal/config"
	"github.com/kayki-araujo/lynk-api/internal/database"
	"github.com/kayki-araujo/lynk-api/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("no .env file found")
	}

	cfg := config.Load()

	db, err := database.Connect(cfg.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("database connected")

	r := router.Setup(db)

	log.Printf("server running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
