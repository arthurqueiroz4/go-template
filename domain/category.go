package domain

type Category struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string
}

type CategoryRepository interface {
	Save(*Category) (*Category, error)
	FindAll(int, int) ([]Category, error)
	FindByID(int) (*Category, error)
	Update(*Category) (*Category, error)
	Delete(int) error
}

type CategoryService interface {
	GetAll(int, int) ([]Category, error)
	GetById(int) (*Category, error)
	Create(*Category) (*Category, error)
	Update(int, *Category) error
	Delete(int) error
}

func NewCategory(id uint, name string) *Category {
	return &Category{
		Name: name,
		ID:   id,
	}
}
