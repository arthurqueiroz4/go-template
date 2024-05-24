package controller

import (
	"crud-golang/api/dto"
	"crud-golang/api/middleware"
	"crud-golang/domain"
	mockdomain "crud-golang/mock"
	"encoding/json"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type testSetup struct {
	ctrl        *gomock.Controller
	mockService *mockdomain.MockCategoryService
	app         *fiber.App
	controller  *CategoryController
}

func setup(t *testing.T) *testSetup {
	ctrl := gomock.NewController(t)
	mockService := mockdomain.NewMockCategoryService(ctrl)
	cc := NewCategoryController(mockService)

	app := fiber.New()
	app.Get("categories/", cc.GetAllCategory)
	app.Get("categories/active", cc.GetAllActive)
	app.Get("categories/:id", cc.GetCategory)
	app.Post("categories/", middleware.ValidationBody[dto.CategoryDTO], cc.Create)
	app.Delete("categories/:id", cc.DeleteCategory)
	app.Put("categories/:id", middleware.ValidationBody[dto.CategoryDTO], cc.UpdateCategory)
	app.Patch("categories/active/:id", middleware.ValidationBody[dto.CategoryDTOActive], cc.UpdateActiveById)

	return &testSetup{
		ctrl:        ctrl,
		mockService: mockService,
		app:         app,
		controller:  cc,
	}
}

func teardown(ts *testSetup) {
	ts.ctrl.Finish()
}

func TestCategoryController_Create(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	categoryDTO := `{"name":"test"}`

	ts.mockService.
		EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(category *domain.Category) error {
			category.ID = 1
			return nil
		})

	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(categoryDTO))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestCategoryController_CreateWithBodyInvalid(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	categoryDTO := `{"active":"name", "name":"test"}`

	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(categoryDTO))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var responseBody map[string]interface{}

	err := json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.NotNil(t, responseBody)
	assert.Equal(t, "body parse error", responseBody["error"])
}

func TestCategoryController_GetCategory(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	category := domain.NewCategory(uint(1), "test", false)

	ts.mockService.
		EXPECT().
		GetById(1).
		Return(category, nil)

	req := httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCategoryController_DeleteCategory(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	ts.mockService.
		EXPECT().
		Delete(1).
		Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/categories/1", nil)
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestCategoryController_UpdateCategory(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	categoryDTO := `{"name":"updated test"}`

	ts.mockService.
		EXPECT().
		Update(1, gomock.Any()).
		Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/categories/1", strings.NewReader(categoryDTO))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCategoryController_GetAllCategory(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	categories := []domain.Category{
		*domain.NewCategory(uint(1), "test1", false),
		*domain.NewCategory(uint(2), "test2", false),
		*domain.NewCategory(uint(3), "test3", false),
	}

	ts.mockService.
		EXPECT().
		GetAll(0, 10, "").
		Return(categories, nil)

	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCategoryController_GetAllActive(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	categories := []domain.Category{
		*domain.NewCategory(uint(1), "test1", true),
		*domain.NewCategory(uint(2), "test2", true),
		*domain.NewCategory(uint(3), "test3", true),
	}

	ts.mockService.
		EXPECT().
		GetAllActive(0, 10).
		Return(categories, nil)

	req := httptest.NewRequest(http.MethodGet, "/categories/active", nil)
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCategoryController_UpdateActiveById(t *testing.T) {
	ts := setup(t)
	defer teardown(ts)

	categoryDTOActive := `{"active":true}`
	category := domain.NewCategory(uint(1), "test", true)

	ts.mockService.
		EXPECT().
		UpdateActive(1, true).
		Return(category, nil)

	req := httptest.NewRequest(http.MethodPatch, "/categories/active/1", strings.NewReader(categoryDTOActive))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := ts.app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
