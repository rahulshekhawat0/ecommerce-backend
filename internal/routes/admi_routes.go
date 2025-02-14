package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/handlers"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/middleware"
)

func SetupAdminRoutes(app *fiber.App) {
	adminRoutes := app.Group("/ecom/admin")

	// Admin-only routes
	adminRoutes.Get("/orders", middleware.JWTMiddleware, handlers.GetAllOrders) // Fetch all orders
}
