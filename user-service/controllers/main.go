package controllers

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"gorm.io/gorm"
)

var (
	DBClient          *gorm.DB
	UserTopicProducer *kafka.Producer
	UserTopic         = "users"
)

func SetDbClient(db *gorm.DB) {
	DBClient = db
}

func SetUserTopicProducer(kw *kafka.Producer) {
	UserTopicProducer = kw
}
