package models

import "time"

// Order represents an order made by a user.
type Order struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `json:"user_id"`                         // Who placed the order
	Status        string    `gorm:"default:'Pending'" json:"status"` // Pending, Shipped, etc.
	TotalPrice    float64   `json:"total_price"`
	PaymentStatus string    `gorm:"default:'Unpaid'" json:"payment_status"` // Unpaid, Paid
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// OrderItem represents a product in an order.
type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"`   // Which order it belongs to
	ProductID uint      `json:"product_id"` // Which product was ordered
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"` // Price at order time
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
