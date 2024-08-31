package controllers

import "gorm.io/gorm"

var (
	DBClient *gorm.DB
)

func SetDbClient(db *gorm.DB) {
	DBClient = db
}
