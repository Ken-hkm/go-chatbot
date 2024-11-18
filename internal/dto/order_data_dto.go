package dto

import "time"

type OrderDataDto struct {
	OrderID         uint         `json:"order_id"`
	SalesNo         string       `json:"sales_no"`
	UserID          uint         `json:"user_id"`
	OrderTotal      float64      `json:"order_total"`
	OrderStatus     string       `json:"order_status"`
	ShippingAddress string       `json:"shipping_address"`
	BillingAddress  string       `json:"billing_address"`
	PaymentMethod   string       `json:"payment_method"`
	PaymentTotal    float64      `json:"payment_total"`
	PaymentStatus   string       `json:"payment_status"`
	TransactionID   string       `json:"transaction_id"`
	PaymentDate     time.Time    `json:"payment_date"`
	Products        []ProductDTO `json:"products"`
}
