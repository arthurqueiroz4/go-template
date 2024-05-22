package main

import (
	"crud-golang/api/route"
	"crud-golang/bootstrap"
	"github.com/gofiber/fiber/v2"
	"log"
)

// TODO criar teste de unidade e integracao
// TODO melhorar erros e padronizar
// TODO dinamic query

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Postgres
	defer app.CloseDBConnection()

	appFiber := fiber.New()

	route.Setup(db, appFiber)

	log.Fatal(appFiber.Listen(env.WebServerPort))
}
