package repository

import (
	"gorm.io/gorm"
)

type PromoRepository interface {
}

type promoRepository struct {
	db *gorm.DB
}

func newPromoRepository(db *gorm.DB) PromoRepository {
	return &promoRepository{db: db}
}
