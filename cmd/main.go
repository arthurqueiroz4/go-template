package main

import (
	"crud-golang/api/route"
	"crud-golang/bootstrap"
	"github.com/gofiber/fiber/v2"
	"log"
)

/*
TODO write readme with explanation about:
	"How gen mocks for tests? < mockgen -source=category.go -destination=./mock/mock_category.go >",
	"Its need that to generate swagger.json with SWAG CLI
	always when something changes for let swagger documentation updated."
*/

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Postgres
	defer app.CloseDBConnection()

	appFiber := fiber.New()

	route.Setup(db, appFiber)

	log.Fatal(appFiber.Listen(env.WebServerPort))
}
