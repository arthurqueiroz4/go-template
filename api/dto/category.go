package dto

import "crud-golang/internal/entity"

type CategoryDTO struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (d CategoryDTO) ParseToEntity() *entity.Category {
	return entity.NewCategory(d.ID, d.Name)
}

func FromEntity(category entity.Category) *CategoryDTO {
	return &CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
}
