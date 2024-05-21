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
