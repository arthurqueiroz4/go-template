package repository

import (
	"crud-golang/domain"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	DB *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return BaseRepository[T]{DB: db}
}

func (b *BaseRepository[T]) Create(category *T) error {
	if err := b.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository[T]) FindAll(page, limit int, params ...string) ([]T, error) {
	var entity []T
	query := b.DB

	for _, param := range params {
		query = query.Where(param)
	}

	if err := query.Limit(limit).Offset(page * limit).Find(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (b *BaseRepository[T]) FindByID(id int) (*T, error) {
	var entity T
	if err := b.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (b *BaseRepository[T]) Update(entity *T) error {
	if err := b.DB.Save(entity).Error; err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository[T]) Delete(id int) error {
	if err := b.DB.Delete(&domain.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
