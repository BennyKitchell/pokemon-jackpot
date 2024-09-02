package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"pokemon-service/pkg/models"
	"strconv"
	"syscall"
	"time"
)

func StartJackpotConsumer() {
	log.Printf("Starting the jackpot consumer")

	// Infinite loop to keep the consumer alivesigChan := make(chan os.Signal, 1)
	JackpotTopicConsumer.SubscribeTopics([]string{JackpotTopic}, nil)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received termination signal, stopping the jackpot consumer.")
		os.Exit(1)
	}()
	for {
		// Check if context is cancelled
		select {
		case <-ctxBackground.Done():
			log.Println("Jackpot Consumer stopped")
			return
		default:
		}

		msg, err := JackpotTopicConsumer.ReadMessage(100 * time.Millisecond)
		if err != nil {
			continue
		}

		userJackpot := models.User_Pokemon{}
		if err := json.Unmarshal(msg.Value, &userJackpot); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		// TODO: Clean this up, not the best way to handle this error
		// Notes: Comparing error to expected sql.ErrNoRows doesn't provide desired behavior
		if DBClient.Where("user_id = ? AND pokemon_id = ?", userJackpot.UserId, userJackpot.PokemonId).First(&userJackpot).Error == nil {
			println("User already owns this pokemon, skipping db update")
			return
		}

		if DBClient.Create(&userJackpot).Error != nil {
			log.Fatal(err)
			return
		}

		pokemon := models.Pokemon{}
		if DBClient.Where("id = ?", userJackpot.UserId).Find(&pokemon).Error != nil {
			log.Printf("Error fetching collection in consumer: %v", err)
			return
		}
		// Update collection in Redis
		userString := strconv.Itoa(userJackpot.UserId)
		redisKey := fmt.Sprintf("collection-%s", userString)
		RedisClient.Set(ctxBackground, redisKey, pokemon, 1*time.Minute)
	}
}
