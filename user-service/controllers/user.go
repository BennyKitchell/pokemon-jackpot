package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"user-service/pkg/models"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		log.Panic(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Clean this up, not the best way to handle this error
	// Notes: Comparing error to expected sql.ErrNoRows doesn't provide desired behavior
	if DBClient.Where("email = ?", user.Email).First(&user).Error == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists!"})
		return
	}

	dbError := DBClient.Create(&user).Error
	if dbError != nil {
		log.Fatal(dbError.Error())
	}
	key := fmt.Sprintf("user-%s", user.Email)

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Panic(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	_, err = RedisClient.Set(context, key, userJson, 1*time.Hour).Result()
	if err != nil {
		log.Fatal(err)
		return
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &UserTopic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(userJson),
	}

	err = UserTopicProducer.Produce(message, nil)

	if err != nil {
		log.Printf("Error writing message to Kafka: %s", err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	// flush the broker, not sure if needed
	go UserTopicProducer.Flush(15 * 1000)

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
