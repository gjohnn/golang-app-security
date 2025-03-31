package database

import (
	"fmt"
	"log"
	"os"
	"v0/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables with fallback error handling
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	postgresPort := os.Getenv("POSTGRES_PORT")

	if postgresHost == "" || postgresUser == "" || postgresPassword == "" || postgresDB == "" || postgresPort == "" {
		return nil, fmt.Errorf("missing one or more required environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		postgresHost,
		postgresUser,
		postgresPassword,
		postgresDB,
		postgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	modelsToMigrate := []interface{}{
		&models.User{},
	}

	// Run migrations for all models in modelsToMigrate
	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			return nil, err
		}
	}

	return db, nil
}
