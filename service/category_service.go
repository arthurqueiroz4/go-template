package service

import (
	"crud-golang/domain"
	"crud-golang/exception"
	"errors"

	"gorm.io/gorm"
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
		return exception.NewErrInternalServer(err.Error(), "error in create category")
	}

	return nil
}

func (cs *CategoryService) GetById(id int) (*domain.Category, error) {
	category, err := cs.cr.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exception.NewErrNotFound(err.Error(), "category not found")
		}
		return nil, exception.NewErrInternalServer(err.Error(), "error in get category by id")
	}

	return category, nil
}

func (cs *CategoryService) GetAll(page, size int, name string) ([]domain.Category, error) {
	all, err := cs.cr.FindAll(page, size, name)
	if err != nil {
		return nil, exception.NewErrInternalServer(err.Error(), "error in get all categories")
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
		return exception.NewErrInternalServer(err.Error(), "error in update category")
	}

	return nil
}

func (cs *CategoryService) Delete(id int) error {
	if _, err := cs.GetById(id); err != nil {
		return err
	}

	if err := cs.cr.Delete(id); err != nil {
		return exception.NewErrInternalServer(err.Error(), "error in delete category")
	}

	return nil
}

func (cs *CategoryService) GetAllActive(page, size int) ([]domain.Category, error) {
	queryForActive := "active = true"

	all, err := cs.cr.FindAll(page, size, queryForActive)
	if err != nil {
		return nil, errors.New("category list failed")
	}

	return all, nil
}

func (cs *CategoryService) UpdateActive(id int, active bool) (*domain.Category, error) {
	categoryToUpdate, err := cs.GetById(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	categoryToUpdate.ID = uint(id)
	categoryToUpdate.Active = active
	err = cs.cr.Update(categoryToUpdate)
	if err != nil {
		return nil, errors.New("category update active failed")
	}

	return categoryToUpdate, nil
}
