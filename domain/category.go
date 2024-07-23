package domain

type Category struct {
	Name     string
	Products []*Product `gorm:"many2many:category_product"`
	ID       uint       `gorm:"primary_key;auto_increment"`
	Active   bool
}

type CategoryRepository interface {
	BaseRepository[Category]
}

type CategoryService interface {
	GetAllActive(int, int) ([]Category, error)
	GetAll(int, int, string) ([]Category, error)
	GetById(int) (*Category, error)
	Create(*Category) error
	Update(int, *Category) error
	UpdateActive(int, bool) (*Category, error)
	Delete(int) error
}

func NewCategory(id uint, name string, active bool) *Category {
	return &Category{
		Name:   name,
		ID:     id,
		Active: active,
	}
}
