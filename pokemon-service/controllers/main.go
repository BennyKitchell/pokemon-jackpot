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
	JackpotTopicConsumer *kafka.Consumer
	JackpotTopicProducer *kafka.Producer
	ctxBackground        = context.Background()
	UserTopic            = "users"
	JackpotTopic         = "jackpot"
)

func SetRedis(rc *redis.Client) {
	RedisClient = rc
}

func SetDbClient(db *gorm.DB) {
	DBClient = db
}

func SetUserTopicConsumer(kc *kafka.Consumer) {
	UserTopicConsumer = kc
}

func SetJackpotTopicConsumer(kc *kafka.Consumer) {
	JackpotTopicConsumer = kc
}

func SetJackpotProducer(kp *kafka.Producer) {
	JackpotTopicProducer = kp
}
