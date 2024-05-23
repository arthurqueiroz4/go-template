package main

import (
	"crud-golang/api/route"
	"crud-golang/bootstrap"
	"github.com/gofiber/fiber/v2"
	"log"
)

//TODO use automapper https://stackoverflow.com/questions/65584357/mapping-one-type-to-another
/*
TODO write readme with explanation about:
	"How to run tests and generate coverage.html?",
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
