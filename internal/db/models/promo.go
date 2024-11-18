package models

// Promo represents promotion details for an order and its products
type Promo struct {
	ID               uint    `gorm:"primaryKey"`
	OrderID          uint    `gorm:"not null"`
	ProductID        uint    `gorm:"not null"`
	PromoCode        string  `gorm:"size:50"`
	Discount         float64 `gorm:"type:decimal(5,2);default:0"`
	PromoDescription string  `gorm:"type:text"`
	Order            Order   `gorm:"foreignKey:OrderID"`
	Product          Product `gorm:"foreignKey:ProductID"`
}
