package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/config"
)

func main() {
	config.ConnectDatabase()

	// Get underlying sql.DB to properly close it
	sqlDB, err := config.DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()

	// Create a new Fiber instance
	app := fiber.New()

	// Define a basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to E-commerce API!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
