package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

func Setup(db gorm.DB, app *fiber.App) {

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "America/Sao_Paulo",
	}))

	publicRouter := app.Group("")
	//privateRouter := app.Group("")

	NewCategoryRouter(db, publicRouter)
}
