package repository

import (
	"errors"
	"go-chatbot/internal/dto"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderDataByUserId(user_id string) ([]dto.OrderDataDto, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetOrderDataByUserId(user_id string) ([]dto.OrderDataDto, error) {
	var orders []dto.OrderDataDto
	// Use raw SQL query for flexibility
	err := r.db.Raw("SELECT o.id as order_id, o.sales_no, o.user_id, o.total_amount AS order_total, o.status AS order_status, o.shipping_address, o.billing_address, p.payment_method, p.total_amount AS payment_total, p.payment_status, p.transaction_id, p.payment_date FROM orders o JOIN payments p ON o.id = p.order_id WHERE o.user_id = ?", user_id).Scan(&orders)
	if err != nil {
		return nil, errors.New("error fetching order data")
	}
	for orderIndex := range orders {
		var products []dto.ProductDTO
		err := r.db.Raw(`SELECT op.order_id, op.product_id, pr.product_name, op.quantity, pr.price, (op.quantity * pr.price) AS total_price, pm.promo_code, pm.discount, pm.promo_description FROM order_products op JOIN products pr ON op.product_id = pr.id LEFT JOIN promo pm ON op.order_id = pm.order_id AND op.product_id = pm.product_id WHERE op.order_id = ?;`, orders[orderIndex].OrderID).Scan(&products).Error
		if err != nil {
			return nil, errors.New("error fetching product data")
		}
		orders[orderIndex].Products = products
	}
	return orders, nil
}
