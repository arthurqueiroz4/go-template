package database

import (
	"crud-golang/internal/entity"
	"gorm.io/gorm"
)

type CategoryImpl struct {
	DB *gorm.DB
}

func NewCategoryDB(db *gorm.DB) *CategoryImpl {
	return &CategoryImpl{
		DB: db,
	}
}

func (c *CategoryImpl) Save(category *entity.Category) (*entity.Category, error) {
	if err := c.DB.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryImpl) FindAll(page, limit int) ([]entity.Category, error) {
	var categories []entity.Category
	if err := c.DB.Limit(limit).Offset(page * limit).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryImpl) FindByID(id uint) (*entity.Category, error) {
	var category entity.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryImpl) Update(category *entity.Category) (*entity.Category, error) {
	if err := c.DB.Save(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryImpl) Delete(id uint) error {
	if err := c.DB.Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
