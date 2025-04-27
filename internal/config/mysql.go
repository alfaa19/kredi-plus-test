package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

func ConnectDB(cfg *Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
	return db
}

func AutoMigrate(db *sql.DB) {
	migrationsDir := "./migrations"

	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		log.Println("No migration folder found, skipping auto-migration...")
		return
	}

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatalf("goose: failed to set dialect: %v", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("goose: migration failed: %v", err)
	}

	fmt.Println("✅ Database migrated successfully")
}
