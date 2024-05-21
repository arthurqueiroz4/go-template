package bootstrap

import (
	"context"
	"crud-golang/domain"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func NewPostgresDatabase(env *Env) gorm.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPassword := env.DBPassword
	dbName := env.DBName

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(errors.New("failed to connection to database"))
	}

	sqlDB, _ := db.DB()
	if err := sqlDB.PingContext(ctx); err != nil {
		panic(errors.New("failed to ping database"))
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&domain.Category{})

	ctx.Done()

	log.Println("Database connection established")
	return *db
}

func ClosePostgresDatabase(db gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connection to database closed.")
}
