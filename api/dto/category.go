package dto

import (
	"crud-golang/domain"
)

type CategoryDTO struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" validate:"required,max=50"`
	Active bool   `json:"active" validate:"boolean"`
}

func (d CategoryDTO) ParseToEntity() *domain.Category {
	return domain.NewCategory(d.ID, d.Name, d.Active)
}

func FromEntity(category domain.Category) *CategoryDTO {
	return &CategoryDTO{
		ID:     category.ID,
		Name:   category.Name,
		Active: category.Active,
	}
}

func FromEntities(categories []domain.Category) []CategoryDTO {
	categoriesDTO := make([]CategoryDTO, len(categories))
	for i, category := range categories {
		categoriesDTO[i] = CategoryDTO{
			ID:     category.ID,
			Name:   category.Name,
			Active: category.Active,
		}
	}
	return categoriesDTO
}
