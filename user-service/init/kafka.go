package initializers

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ConnectProducerToKafka() (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":       "kafka:9093",
		"api.version.fallback.ms": 0})
}
