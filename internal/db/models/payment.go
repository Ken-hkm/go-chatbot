package models

import "time"

// Payment represents the payment information for an order
type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	OrderID       uint      `gorm:"not null"`
	TransactionID string    `gorm:"size:100;not null"`
	PaymentMethod string    `gorm:"size:50;not null"`
	PaymentDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	TotalAmount   float64   `gorm:"not null"`
	PaymentStatus string    `gorm:"size:20;default:'Pending'"`
	Order         Order     `gorm:"foreignKey:OrderID"`
}
