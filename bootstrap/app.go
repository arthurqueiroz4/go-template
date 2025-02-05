package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env      *Env
	Postgres gorm.DB
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	return app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDatabase(app.Postgres)
}
