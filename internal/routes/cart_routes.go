package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/handlers"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/middleware"
)

func SetupCartRoutes(app *fiber.App) {
	cart := app.Group("/ecom/cart")
	// protect these routes with JWT middleware.
	cart.Post("/", middleware.JWTMiddleware, handlers.AddCartItem)
	//  assuming the user ID is passed as a URL parameter.
	cart.Get("/:userId", middleware.JWTMiddleware, handlers.GetCartItems)
	cart.Put("/:id", middleware.JWTMiddleware, handlers.UpdateCartItem)
	cart.Delete("/:id", middleware.JWTMiddleware, handlers.DeleteCartItem)
}
