package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/alfaa19/kredi-plus-test/internal/app"
	"github.com/alfaa19/kredi-plus-test/internal/config"
)

func main() {
	// Load .env manually
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, skipping...")
	}

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect DB
	db := config.ConnectDB(cfg)
	defer db.Close()

	config.AutoMigrate(db)

	application := app.NewApp(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServerPort
	}

	log.Println("🚀 Starting server on port", port)
	if err := application.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
