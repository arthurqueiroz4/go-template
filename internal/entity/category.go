package entity

type Category struct {
	ID   uint `gorm:"primary_key;auto_increment"`
	Name string
}

func NewCategory(id uint, name string) *Category {
	return &Category{
		Name: name,
		ID:   id,
	}
}
