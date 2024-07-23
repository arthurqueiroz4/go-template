package dto

type CategoryDTO struct {
	Name     string       `json:"name" validate:"required,max=50"`
	Products []ProductDTO `json:"products"`
	ID       uint         `json:"id"`
	Active   bool         `json:"active" validate:"boolean"`
}

type CategoryDTOActive struct {
	Active bool `json:"active" validate:"boolean"`
}
