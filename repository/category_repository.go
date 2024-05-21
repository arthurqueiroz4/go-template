package repository

import (
	"crud-golang/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryDB(db *gorm.DB) domain.CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (c *CategoryRepositoryImpl) Save(category *domain.Category) (*domain.Category, error) {
	if err := c.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) FindAll(page, limit int) ([]domain.Category, error) {
	var categories []domain.Category
	if err := c.db.Limit(limit).Offset(page * limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryRepositoryImpl) FindByID(id int) (*domain.Category, error) {
	var category domain.Category
	if err := c.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepositoryImpl) Update(category *domain.Category) (*domain.Category, error) {
	if err := c.db.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) Delete(id int) error {
	if err := c.db.Delete(&domain.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
