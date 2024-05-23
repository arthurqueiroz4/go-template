package domain

type (
	BaseRepository[T any] interface {
		Create(*T) error
		FindAll(int, int, ...string) ([]T, error)
		FindByID(int) (*T, error)
		Update(*T) error
		Delete(int) error
	}
)
