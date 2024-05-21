package main

import (
	"crud-golang/api/route"
	"crud-golang/config"
	"crud-golang/internal/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Category{})
	if err != nil {
		return
	}

	route.Setup(db, cfg)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
