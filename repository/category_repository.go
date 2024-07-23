package repository

import (
	"crud-golang/domain"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	BaseRepository[domain.Category]
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{NewBaseRepository[domain.Category](db)}
}
