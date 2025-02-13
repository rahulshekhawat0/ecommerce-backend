package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB Global variable
var DB *gorm.DB

// ConnectDatabase initializes DB connection
func ConnectDatabase() *gorm.DB {
	// Load .env file (if available)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	// Required for Neon

	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSLMode,
	)

	// Open GORM database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database Connection Failed: %v", err)
	}

	log.Println("Connected to Database successfully!")

	// Assign DB to global variable
	DB = db

	// AutoMigrate Models (Add models as needed)
	if err := db.AutoMigrate(&models.User{},
		&models.Product{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{}); err != nil {
		log.Fatalf("Migration Failed: %v", err)
	}

	log.Println(" Database Migrated Successfully!")
	return db
}
