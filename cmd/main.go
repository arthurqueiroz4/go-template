package main

import (
	"crud-golang/api/route"
	"crud-golang/bootstrap"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Postgres
	defer app.CloseDBConnection()

	appFiber := fiber.New()

	route.Setup(db, appFiber)

	log.Fatal(appFiber.Listen(env.WebServerPort))
}
