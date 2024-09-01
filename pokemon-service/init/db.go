package initializers

import (
	"context"
	"log"
	db_connection "pokemon-service/modules/db"
	"pokemon-service/pkg/models"
	"time"

	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db := db_connection.NewDBClient(context.Background())
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	db.AutoMigrate(models.Pokemon{})      // Migrate the Pokemon model
	db.AutoMigrate(models.User_Pokemon{}) // Migrate the User_Pokemon model

	return db
}
