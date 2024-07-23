package repository

import (
	"crud-golang/domain"

	"gorm.io/gorm"
)

type ProductRepository struct {
	BaseRepository[domain.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{NewBaseRepository[domain.Product](db)}
}
