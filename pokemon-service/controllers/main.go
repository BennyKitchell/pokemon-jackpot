package controllers

import (
	"context"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	RedisClient          *redis.Client
	DBClient             *gorm.DB
	UserTopicConsumer    *kafka.Consumer
	ctxBackground        = context.Background()
	UserTopic            = "users"
	JackpotTopic         = "jackpot"
	JackpotTopicProducer *kafka.Producer
)

func SetRedis(rc *redis.Client) {
	RedisClient = rc
}

func SetDbClient(db *gorm.DB) {
	DBClient = db
}

func SetUserTopicConsumer(kr *kafka.Consumer) {
	UserTopicConsumer = kr
}

func SetJackpotProducer(kp *kafka.Producer) {
	JackpotTopicProducer = kp
}
