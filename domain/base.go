package domain

type (
	BaseRepository[T any] interface {
		Create(*T) error
		FindAll(int, int) ([]T, error)
		FindByID(int) (*T, error)
		Update(*T) error
		Delete(int) error
	}
)
