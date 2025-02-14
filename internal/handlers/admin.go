package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/config"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/models"
)

// GetAllOrders allows admins to fetch all orders from all users.
func GetAllOrders(c *fiber.Ctx) error {
	// Ensure the user is an admin (Extract role from JWT)
	userRole, ok := c.Locals("userRole").(string)
	if !ok || userRole != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "Access denied"})
	}

	//Fetch all orders, including order items
	var orders []models.Order
	if err := config.DB.Preload("Items").Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch orders"})
	}

	// Return the orders in JSON format
	return c.JSON(fiber.Map{"orders": orders})
}
