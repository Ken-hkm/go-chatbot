package repository

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
}

type productRepository struct {
	db *gorm.DB
}

func newProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}
