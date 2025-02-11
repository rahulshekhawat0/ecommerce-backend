package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var DB *sql.DB

func ConnectDatabase() {
	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE") // Required for Neon

	// Create DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSLMode,
	)
	// Open database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error connecting to the database:", err)
	}
	// Ping the database to check connection
	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Database is not reachable:", err)
	}
	fmt.Println("✅ Connected to Database successfully!")
	// Assign db to global variable
	DB = db
}
