package dto

type ProductDTO struct {
	OrderID          uint    `json:"order_id"`
	ProductID        uint    `json:"product_id"`
	ProductName      string  `json:"product_name"`
	Quantity         int     `json:"quantity"`
	Price            float64 `json:"price"`
	TotalPrice       float64 `json:"total_price"`
	PromoCode        string  `json:"promo_code,omitempty"`
	Discount         float64 `json:"discount,omitempty"`
	PromoDescription string  `json:"promo_description,omitempty"`
}
