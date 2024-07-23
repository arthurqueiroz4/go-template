package controller

import (
	"crud-golang/api/dto"
	"crud-golang/domain"

	"github.com/PeteProgrammer/go-automapper"
	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	cs domain.CategoryService
}

func NewCategoryController(cs domain.CategoryService) *CategoryController {
	return &CategoryController{
		cs: cs,
	}
}

// Create Category
//
// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body dto.CategoryDTO true "Category DTO"
// @Success 200 {object} dto.CategoryDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /categories [post]
func (cc *CategoryController) Create(c *fiber.Ctx) error {
	var categoryDTO dto.CategoryDTO

	if err := c.BodyParser(&categoryDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "Invalid request body"})
	}

	var category domain.Category
	automapper.MapLoose(categoryDTO, &category)

	err := cc.cs.Create(&category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{"message": "Could not create category"})
	}

	automapper.Map(category, &categoryDTO)

	return c.Status(fiber.StatusCreated).
		JSON(categoryDTO)
}

// GetCategory
//
// @Summary Get a category by ID
// @Description Retrieve a category by its ID
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} dto.CategoryDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /categories/{id} [get]
func (cc *CategoryController) GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "ID is required"})
	}
	category, err := cc.cs.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(map[string]any{"message": "Category not found"})
	}

	var categoryDTO dto.CategoryDTO
	automapper.Map(category, &categoryDTO)
	return c.Status(fiber.StatusOK).
		JSON(categoryDTO)
}

// DeleteCategory
//
// @Summary Delete a category by ID
// @Description Delete a category by its ID
// @Tags Categories
// @Param id path int true "Category ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /categories/{id} [delete]
func (cc *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "ID is required"})
	}

	err = cc.cs.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(map[string]any{"message": "Category not found"})
	}

	c.Status(fiber.StatusNoContent)
	return nil
}

// UpdateCategory
//
// @Summary Update a category by ID
// @Description Update a category with the provided details
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body dto.CategoryDTO true "Category DTO"
// @Success 200 {object} dto.CategoryDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /categories/{id} [put]
func (cc *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "ID is required"})
	}

	var categoryDTO dto.CategoryDTO
	if err := c.BodyParser(&categoryDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "Invalid request body"})
	}

	var category domain.Category
	automapper.Map(categoryDTO, &category)
	err = cc.cs.Update(id, &category)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(map[string]any{"message": "Category not found"})
	}

	automapper.Map(category, &categoryDTO)
	return c.Status(fiber.StatusOK).
		JSON(categoryDTO)
}

// GetAllCategory
//
// @Summary Get all categories with pagination and optional filtering by name
// @Description Retrieve all categories, with pagination and optional filtering by name
// @Tags Categories
// @Produce json
// @Param page query int false "Page number" default(0)
// @Param size query int false "Page size" default(10)
// @Param name query string false "Category name"
// @Success 200 {array} dto.CategoryDTO
// @Failure 500 {object} map[string]interface{}
// @Router /categories [get]
func (cc *CategoryController) GetAllCategory(c *fiber.Ctx) error {
	page := c.QueryInt("page", 0)
	size := c.QueryInt("size", 10)
	name := c.Query("name")

	all, err := cc.cs.GetAll(page, size, name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{"message": "Internal server error"})
	}

	var dtos []dto.CategoryDTO
	automapper.Map(all, &dtos)
	return c.Status(fiber.StatusOK).
		JSON(dtos)
}

// GetAllActive
//
// @Summary Get all categories with pagination and filtering by active
// @Description Retrieve all categories, with pagination and filtering by active
// @Tags Categories
// @Produce json
// @Param page query int false "Page number" default(0)
// @Param size query int false "Page size" default(10)
// @Success 200 {array} dto.CategoryDTO
// @Failure 500 {object} map[string]interface{}
// @Router /categories/active [get]
func (cc *CategoryController) GetAllActive(c *fiber.Ctx) error {
	page := c.QueryInt("page", 0)
	size := c.QueryInt("size", 10)

	all, err := cc.cs.GetAllActive(page, size)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{"message": "Internal server error"})
	}

	var dtos []dto.CategoryDTO
	automapper.Map(all, &dtos)
	return c.Status(fiber.StatusOK).
		JSON(dtos)
}

// UpdateActiveById
//
// @Summary Update a category status by ID
// @Description Update a category activation status
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param categoryActive body dto.CategoryDTOActive true "Category DTO Active"
// @Success 200 {object} dto.CategoryDTO
// @Router /categories/active/{id} [patch]
func (cc *CategoryController) UpdateActiveById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "ID is required"})
	}

	var categoryDTOActive dto.CategoryDTOActive
	if err := c.BodyParser(&categoryDTOActive); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "Invalid request body"})
	}

	category, err := cc.cs.UpdateActive(id, categoryDTOActive.Active)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(map[string]any{"message": "Category not found"})
	}

	var categoryDTO dto.CategoryDTO
	automapper.Map(category, &categoryDTO)
	return c.Status(fiber.StatusOK).
		JSON(categoryDTO)
}
