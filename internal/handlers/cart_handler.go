package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/config"
	"github.com/rahulshekhawat0/ecommerce-backend/internal/models"
)

// AddCartItem adds a new item to the user's cart.
func AddCartItem(c *fiber.Ctx) error {
	var cartItem models.CartItem
	if err := c.BodyParser(&cartItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Optionally: Validate if product exists and if quantity is valid.

	if err := config.DB.Create(&cartItem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not add item to cart"})
	}
	return c.Status(fiber.StatusCreated).JSON(cartItem)
}

// GetCartItems retrieves all cart items for a user.
func GetCartItems(c *fiber.Ctx) error {
	userIDParam := c.Params("userId")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var cartItems []models.CartItem
	if err := config.DB.Where("user_id = ?", userID).Find(&cartItems).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch cart items"})
	}
	return c.JSON(cartItems)
}

// UpdateCartItem updates the quantity of an existing cart item.
func UpdateCartItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cart item ID"})
	}

	var cartItem models.CartItem
	if err := config.DB.First(&cartItem, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Cart item not found"})
	}

	if err := c.BodyParser(&cartItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	config.DB.Save(&cartItem)
	return c.JSON(cartItem)
}

// DeleteCartItem removes an item from the cart.
func DeleteCartItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid cart item ID"})
	}

	if err := config.DB.Delete(&models.CartItem{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cart item"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
