package domain

type Category struct {
	ID     uint `gorm:"primary_key;auto_increment"`
	Name   string
	Active bool
}

type CategoryRepository interface {
	BaseRepository[Category]
	FindAllSpec(int, int, string) ([]Category, error)
}

type CategoryService interface {
	GetAllActive(int, int) ([]Category, error)
	GetAll(int, int, string) ([]Category, error)
	GetById(int) (*Category, error)
	Create(*Category) error
	Update(int, *Category) error
	Delete(int) error
}

func NewCategory(id uint, name string, active bool) *Category {
	return &Category{
		Name:   name,
		ID:     id,
		Active: active,
	}
}
