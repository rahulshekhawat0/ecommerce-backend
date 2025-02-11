package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Define a basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to E-commerce API!")
	})

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
