package controllers

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	DBClient          *gorm.DB
	UserTopicProducer *kafka.Producer
	RedisClient       *redis.Client
	UserTopic         = "users"
)

func SetDbClient(db *gorm.DB) {
	DBClient = db
}

func SetUserTopicProducer(kw *kafka.Producer) {
	UserTopicProducer = kw
}

func SetRedis(rc *redis.Client) {
	RedisClient = rc
}
