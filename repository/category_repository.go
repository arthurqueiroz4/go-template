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

func (c *CategoryRepository) FindAllSpec(page, size int, name string) ([]domain.Category, error) {
	var categories []domain.Category
	query := c.DB
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query = query.Limit(size).Offset(page * size)

	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
