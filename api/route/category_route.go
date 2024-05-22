package route

import (
	"crud-golang/api/controller"
	"crud-golang/repository"
	"crud-golang/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewCategoryRouter(db gorm.DB, r fiber.Router) {
	cdb := repository.NewCategoryRepository(&db)
	cs := service.NewCategoryService(cdb)
	cc := controller.NewCategoryController(cs)

	r.Get("categories/", cc.GetAllCategory)
	r.Get("categories/:id", cc.GetCategory)
	r.Post("categories/", cc.Create)
	r.Delete("categories/:id", cc.DeleteCategory)
	r.Put("categories/:id", cc.UpdateCategory)
}
