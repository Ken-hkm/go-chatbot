package models

import "time"

type Order struct {
	ID              uint      `gorm:"primaryKey"`
	SalesNo         string    `gorm:"size:50;not null"`
	UserID          uint      `gorm:"not null"` // Changed from CustomerID to UserID
	OrderDate       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	TotalAmount     float64   `gorm:"not null"`
	Status          string    `gorm:"size:20;default:'Pending'"`
	ShippingAddress string    `gorm:"size:255"`
	BillingAddress  string    `gorm:"size:255"`
	User            User      `gorm:"foreignKey:UserID"` // Relationship with User
	Products        []Product `gorm:"foreignKey:OrderID"`
	Promos          []Promo   `gorm:"foreignKey:OrderID"`
	Payments        []Payment `gorm:"foreignKey:OrderID"`
}
