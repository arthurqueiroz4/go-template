package database

import "crud-golang/internal/entity"

type CategoryRepo interface {
	FindAll(int, int) ([]entity.Category, error)
	FindByID(uint) (*entity.Category, error)
	Save(*entity.Category) (*entity.Category, error)
	Update(*entity.Category) (*entity.Category, error)
	Delete(uint) error
}
