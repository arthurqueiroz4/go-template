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

func (cs *CategoryService) Create(category *domain.Category) error {
	err := cs.cr.Create(category)
	if err != nil {
		return errors.New("category creation failed")
	}

	return nil
}

func (cs *CategoryService) GetById(id int) (*domain.Category, error) {
	category, err := cs.cr.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	return category, nil
}

func (cs *CategoryService) GetAll(page, size int, name string) ([]domain.Category, error) {
	all, err := cs.cr.FindAllSpec(page, size, name)
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
	err := cs.cr.Update(categoryToUpdate)
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

func (cs *CategoryService) GetAllActive(page, size int) ([]domain.Category, error) {
	queryForActive := "active = true"

	all, err := cs.cr.FindAll(page, size, queryForActive)
	if err != nil {
		return nil, errors.New("category list failed")
	}

	return all, nil
}
