package initializers

import (
	"context"
	"log"
	"time"
	db_connection "user-service/modules/db"
	"user-service/pkg/models"

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
	db.AutoMigrate(models.User{}) // Migrate the User model

	return db
}
