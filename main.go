package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/config"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/routes"
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
	routes.SetupAuthRoutes(app)
	routes.SetupProductRoutes(app)
	routes.SetupCartRoutes(app)
	routes.SetupOrderRoutes(app)
	routes.SetupAdminRoutes(app)
	// Define a basic route
	app.Get("/ecom", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to E-commerce API!")
	})
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("error in getting Port Value") // Default to port 8000 if PORT is not set
	}

	// Start the server
	log.Fatal(app.Listen(":" + port))

}
