package dto

type CategoryDTO struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" validate:"required,max=50"`
	Active bool   `json:"active" validate:"boolean"`
}

type CategoryDTOActive struct {
	Active bool `json:"active" validate:"boolean"`
}
