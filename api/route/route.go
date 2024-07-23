package route

import (
	_ "crud-golang/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

// @title Category CRUD
// @version 1.0
// @description This is a sample CRUD.
// @contact.name API Support
// @contact.url http://suport.com
// @contact.email support@dev.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @schemes http
func Setup(db gorm.DB, app *fiber.App) {
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "America/Sao_Paulo",
	}))

	publicRouter := app.Group("")
	publicRouter.Get("/swagger/*", swagger.HandlerDefault)
	// privateRouter := app.Group("")

	NewCategoryRouter(db, publicRouter)
	NewProductRouter(db, publicRouter)
}
