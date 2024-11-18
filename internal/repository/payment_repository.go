package repository

import (
	"gorm.io/gorm"
)

type PaymentRepository interface {
}

type paymentRepository struct {
	db *gorm.DB
}

func newPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}
