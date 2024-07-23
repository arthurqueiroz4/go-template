package controller

import (
	"crud-golang/api/dto"
	"crud-golang/domain"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ps domain.ProductService
}

func NewProductController(ps domain.ProductService) ProductController {
	return ProductController{ps}
}

func (pc *ProductController) Create(c *fiber.Ctx) error {
	var dto dto.ProductDTO

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "Invalid request body"})
	}

	product := domain.Product{Name: dto.Name}

	if pc.ps.Create(&product) != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "Error in product creation"})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}

func (pc *ProductController) GetAll(c *fiber.Ctx) error {
	p, err := pc.ps.GetAll(0, 100, "")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{"message": "Error in product list"})
	}

	return c.Status(fiber.StatusOK).
		JSON(p)
}
