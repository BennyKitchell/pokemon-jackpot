package initializers

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ConnectProducerToKafka() (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092"})
}
