package domain

type Product struct {
	Name       string
	Categories []*Category `gorm:"many2many:category_product"`
	ID         uint        `gorm:"primary_key;auto_increment"`
}

type ProductRepository interface {
	BaseRepository[Product]
}

type ProductService interface {
	GetAll(int, int, string) ([]Product, error)
	// GetById(int) (*Product, error)
	Create(*Product) error
	// Update(int, *Product) error
	// Delete(int) error
}
