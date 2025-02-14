package handlers

import (
	"fmt"
	"strconv"
	"time"

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

func GetOrders(c *fiber.Ctx) error {
	// Extract user ID from JWT
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Fetch all orders for this user, including items
	var orders []models.Order
	if err := config.DB.Preload("Items").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch orders"})
	}

	// Return the orders in JSON format
	return c.Status(200).JSON(fiber.Map{"orders": orders})
}

func CancelOrder(c *fiber.Ctx) error {
	// Extract user ID from JWT
	userID, ok := c.Locals("userID").(uint)
	if !ok {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	//Get order ID from the request params
	orderID := c.Params("id")

	//Find the order
	var order models.Order
	if err := config.DB.First(&order, "id = ? AND user_id = ?", orderID, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}

	//Only allow cancellation if the order is still "Pending"
	if order.Status != "Pending" {
		return c.Status(400).JSON(fiber.Map{"error": "Order cannot be canceled"})
	}

	//Update the order status to "Canceled"
	order.Status = "Canceled"
	if err := config.DB.Save(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to cancel order"})
	}

	return c.JSON(fiber.Map{"message": "Order canceled", "order": order})
}

func UpdateOrderStatus(c *fiber.Ctx) error {
	// Ensure the user is an admin
	userRole, ok := c.Locals("userRole").(string)
	if !ok || userRole != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden: Admins only"})
	}

	// Get order ID from URL params
	orderID := c.Params("id")

	//  Parse the request body (New Status)
	var request struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	//  Validate status (only allow "Shipped" or "Delivered")
	validStatuses := map[string]bool{"Shipped": true, "Delivered": true}
	if !validStatuses[request.Status] {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid status value"})
	}

	//  Find the order by ID
	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}

	// Update the order status
	order.Status = request.Status
	if err := config.DB.Save(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update order status"})
	}

	return c.JSON(fiber.Map{"message": "Order status updated", "order": order})
}

func SimulatePayment(c *fiber.Ctx) error {
	// Get the order ID from the URL params
	orderID := c.Params("id")

	// Find the order
	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}

	// Check if the order is already paid
	if order.PaymentStatus == "Paid" {
		return c.Status(400).JSON(fiber.Map{"error": "Order already paid"})
	}

	// Simulate payment success by updating payment status
	order.PaymentStatus = "Paid"
	config.DB.Save(&order)

	// Generate a fake transaction ID
	transactionID := fmt.Sprintf("TXN%d", time.Now().Unix())

	return c.JSON(fiber.Map{
		"message":        "Payment successful",
		"order_id":       order.ID,
		"total_price":    order.TotalPrice,
		"payment_status": order.PaymentStatus,
		"transaction_id": transactionID,
	})
}
