package route

import (
	"crud-golang/api/controller"
	"crud-golang/api/dto"
	"crud-golang/api/middleware"
	"crud-golang/repository"
	"crud-golang/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewProductRouter(db gorm.DB, r fiber.Router) {
	pr := repository.NewProductRepository(&db)
	ps := service.NewProductService(pr)
	pc := controller.NewProductController(ps)

	r.Get("/products", pc.GetAll)
	r.Post("/products", middleware.ValidationBody[dto.ProductDTO], pc.Create)
}
