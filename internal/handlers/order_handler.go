package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/config"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/models"
)

// Checkout converts all cart items of a user into an order.
func Checkout(c *fiber.Ctx) error {
	// For simplicity, assume user ID is provided in the URL.
	userIDParam := c.Params("userId")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	// Retrieve all cart items for the user.
	var cartItems []models.CartItem
	if err := config.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch cart items"})
	}

	if len(cartItems) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cart is empty"})
	}

	// Calculate total price (using a fixed price of 10.0 for each item as a placeholder).
	totalPrice := 0.0
	for _, item := range cartItems {
		totalPrice += float64(item.Quantity) * 10.0
	}

	// Create the order.
	order := models.Order{
		UserID:     uint(userID),
		Status:     "Pending",
		TotalPrice: totalPrice,
	}
	if err := config.DB.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create order"})
	}

	// Create order items for each cart item.
	for _, item := range cartItems {
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     10.0, // Placeholder price
		}
		config.DB.Create(&orderItem)
	}

	// Clear the user's cart.
	config.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})

	return c.Status(fiber.StatusCreated).JSON(order)
}
