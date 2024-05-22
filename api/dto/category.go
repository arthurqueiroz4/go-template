package dto

import (
	"crud-golang/domain"
)

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (d CategoryDTO) ParseToEntity() *domain.Category {
	return domain.NewCategory(d.ID, d.Name)
}

func FromEntity(category domain.Category) *CategoryDTO {
	return &CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}

func FromEntities(categories []domain.Category) []CategoryDTO {
	categoriesDTO := make([]CategoryDTO, len(categories))
	for i, category := range categories {
		categoriesDTO[i] = CategoryDTO{
			ID:   category.ID,
			Name: category.Name,
		}
	}
	return categoriesDTO
}
