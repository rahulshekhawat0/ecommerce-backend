package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/handlers"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/middleware"
)

func SetupProductRoutes(app *fiber.App) {
	product := app.Group("/ecom/products")
	product.Get("/", handlers.GetProducts)
	product.Get("/:id", handlers.GetProductbyID)
	// Protected routes (assumes JWTMiddleware is implemented)
	product.Post("/", middleware.JWTMiddleware, handlers.CreateProduct)
	product.Put("/:id", middleware.JWTMiddleware, handlers.UpdateProduct)
	product.Delete("/:id", middleware.JWTMiddleware, handlers.DeleteProduct)
}
