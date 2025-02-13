package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/handlers"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/middleware"
)

// SetupOrderRoutes registers the order management endpoints.
func SetupOrderRoutes(app *fiber.App) {
	orders := app.Group("/ecom/orders")
	// Protected route using JWT middleware (for simplicity, the user ID is still passed in the URL)
	orders.Post("/checkout/:userId", middleware.JWTMiddleware, handlers.Checkout)
}
