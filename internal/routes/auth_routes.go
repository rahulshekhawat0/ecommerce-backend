package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/handlers"
)

// SetupAuthRoutes registers authentication routes on the Fiber app
func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/ecom/auth")
	auth.Post("/register", handlers.RegisterUser)
	auth.Post("/login", handlers.LoginUser)
}
