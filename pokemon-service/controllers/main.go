package controllers

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	RedisClient *redis.Client
	DBClient    *gorm.DB
)

func SetRedis(rc *redis.Client) {
	RedisClient = rc
}

func SetDbClient(db *gorm.DB) {
	DBClient = db
}
