package db

import (
	"context"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBClient(context context.Context) *gorm.DB {
	dsn := "postgres://postgres:postgres@localhost:5432?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	return db
}
