package controller

import (
	"crud-golang/api/dto"
	"crud-golang/domain"
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

func (cc *CategoryController) Create(c *fiber.Ctx) error {
	var categoryDTO dto.CategoryDTO

	if err := c.BodyParser(&categoryDTO); err != nil {
		return err
	}

	category := categoryDTO.ParseToEntity()

	err := cc.cs.Create(category)
	if err != nil {
		return err
	}

	return c.JSON(dto.FromEntity(*category))
}

func (cc *CategoryController) GetCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "id is required"})
	}
	category, err := cc.cs.GetById(id)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "category not found"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(dto.FromEntity(*category))
}

func (cc *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "id is required"})
	}

	err = cc.cs.Delete(id)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "category not found"})
	}

	c.Status(fiber.StatusNoContent)
	return nil
}

func (cc *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"message": "id is required"})
	}

	var categoryDTO dto.CategoryDTO
	if err := c.BodyParser(&categoryDTO); err != nil {
		return err
	}
	err = cc.cs.Update(id, categoryDTO.ParseToEntity())
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{"message": "category not found"})
	}

	c.Status(fiber.StatusNoContent)
	return nil
}

func (cc *CategoryController) GetAllCategory(c *fiber.Ctx) error {
	page := c.QueryInt("page", 0)
	size := c.QueryInt("size", 10)
	name := c.Query("name")

	all, err := cc.cs.GetAll(page, size, name)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{"message": "internal server error"})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(all)
}
