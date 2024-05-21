package service

import (
	"crud-golang/domain"
	"errors"
)

type CategoryService struct {
	cr domain.CategoryRepository
}

func NewCategoryService(categoryRepository domain.CategoryRepository) domain.CategoryService {
	return &CategoryService{
		cr: categoryRepository,
	}
}

func (cs *CategoryService) Create(category *domain.Category) (*domain.Category, error) {
	categorySaved, err := cs.cr.Save(category)
	if err != nil {
		return nil, errors.New("category creation failed")
	}

	return categorySaved, nil
}

func (cs *CategoryService) GetById(id int) (*domain.Category, error) {
	category, err := cs.cr.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func (cs *CategoryService) GetAll(page, size int) ([]domain.Category, error) {
	all, err := cs.cr.FindAll(page, size)
	if err != nil {
		return nil, errors.New("category list failed")
	}

	return all, nil
}

func (cs *CategoryService) Update(id int, categoryToUpdate *domain.Category) error {
	if _, err := cs.GetById(id); err != nil {
		return err
	}

	categoryToUpdate.ID = uint(id)
	_, err := cs.cr.Update(categoryToUpdate)
	if err != nil {
		return errors.New("category update failed")
	}

	return nil
}

func (cs *CategoryService) Delete(id int) error {
	if _, err := cs.GetById(id); err != nil {
		return err
	}

	return cs.cr.Delete(id)
}
