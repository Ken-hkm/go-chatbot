package models

// Product represents a product related to an order
type Product struct {
	ID          uint    `gorm:"primaryKey"`
	OrderID     uint    `gorm:"not null"`
	ProductCode string  `gorm:"size:50;not null"`
	ProductName string  `gorm:"size:100;not null"`
	Qty         int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	TotalPrice  float64 `gorm:"->;<-;type:decimal(10,2);"`
	Order       Order   `gorm:"foreignKey:OrderID"`
	Promo       []Promo `gorm:"foreignKey:ProductID"`
}
