package service

import (
	"crud-golang/domain"
	mockdomain "crud-golang/domain/mock"
	"errors"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := mockdomain.NewMockCategoryRepository(ctrl)

	mr.
		EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(category *domain.Category) error {
			category.ID = 1
			return nil
		}).AnyTimes()

	categoryService := NewCategoryService(mr)

	c := domain.Category{
		Name: "test",
	}

	err := categoryService.Create(&c)

	assert.Nil(t, err)
	assert.Equal(t, uint(1), c.ID)
	assert.Equal(t, "test", c.Name)
}

func TestShouldGetAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := mockdomain.NewMockCategoryRepository(ctrl)

	categories := []domain.Category{
		*domain.NewCategory(uint(1), "test1"),
		*domain.NewCategory(uint(2), "test2"),
		*domain.NewCategory(uint(3), "test3"),
	}

	mr.
		EXPECT().
		FindAllSpec(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(categories, nil).
		AnyTimes()

	categoryService := NewCategoryService(mr)

	all, err := categoryService.GetAll(0, 10, "")

	assert.Nil(t, err)
	assert.Equal(t, len(categories), len(all))
	assert.Equal(t, categories, all)
}

func TestShouldGetCategoryById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := mockdomain.NewMockCategoryRepository(ctrl)

	category := domain.NewCategory(uint(1), "test")

	mr.
		EXPECT().
		FindByID(gomock.Any()).
		Return(category, nil).
		AnyTimes()

	categoryService := NewCategoryService(mr)

	c, err := categoryService.GetById(1)

	assert.Nil(t, err)
	assert.Equal(t, category, c)
}

func TestShouldUpdateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := mockdomain.NewMockCategoryRepository(ctrl)

	category := domain.NewCategory(uint(1), "test")
	updatedCategory := domain.NewCategory(uint(1), "updated test")

	mr.
		EXPECT().
		FindByID(gomock.Any()).
		Return(category, nil).
		AnyTimes()

	mr.
		EXPECT().
		Update(gomock.Any()).
		Return(nil).
		AnyTimes()

	categoryService := NewCategoryService(mr)

	err := categoryService.Update(1, updatedCategory)

	assert.Nil(t, err)
	assert.Equal(t, uint(1), updatedCategory.ID)
	assert.Equal(t, "updated test", updatedCategory.Name)
}

func TestShouldDeleteCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := mockdomain.NewMockCategoryRepository(ctrl)

	category := domain.NewCategory(uint(1), "test")

	mr.
		EXPECT().
		FindByID(gomock.Any()).
		Return(category, nil).
		AnyTimes()

	mr.
		EXPECT().
		Delete(gomock.Any()).
		Return(nil).
		AnyTimes()

	categoryService := NewCategoryService(mr)

	err := categoryService.Delete(1)

	assert.Nil(t, err)
}

func TestShouldReturnErrorWhenCategoryNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := mockdomain.NewMockCategoryRepository(ctrl)

	mr.
		EXPECT().
		FindByID(gomock.Any()).
		Return(nil, errors.New("not found")).
		AnyTimes()

	categoryService := NewCategoryService(mr)

	_, err := categoryService.GetById(1)
	assert.NotNil(t, err)
	assert.Equal(t, "category not found", err.Error())

	err = categoryService.Update(1, &domain.Category{Name: "test"})
	assert.NotNil(t, err)
	assert.Equal(t, "category not found", err.Error())

	err = categoryService.Delete(1)
	assert.NotNil(t, err)
	assert.Equal(t, "category not found", err.Error())
}
