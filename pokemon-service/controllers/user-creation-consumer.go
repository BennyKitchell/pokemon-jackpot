package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"os/signal"
	"pokemon-service/pkg/models"
	"syscall"
	"time"
)

func StartUserCreationConsumer() {
	log.Println("Starting the user creation consumer")

	// Infinite loop to keep the consumer alivesigChan := make(chan os.Signal, 1)
	UserTopicConsumer.SubscribeTopics([]string{UserTopic}, nil)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("Received termination signal. Stopping the user creation consumer")
		os.Exit(1)
	}()
	for {
		select {
		case <-ctxBackground.Done():
			log.Println("User Creation Consumer has stopped")
			return
		default:
		}

		msg, err := UserTopicConsumer.ReadMessage(100 * time.Millisecond)
		if err != nil {
			log.Printf("Error reading message from Kafka: %v", err)
			continue
		}

		user := models.User{}
		if err := json.Unmarshal(msg.Value, &user); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		// debug
		fmt.Printf("User Creation Message: %d, email: %s, password: %s, email: %s\n", user.ID, user.Email, user.Password, user.Email)

		// Assign the user one of the starters from the games :easter-egg:
		starterPokemon := [4]int{1, 4, 7, 25}
		UserPokemon := models.User_Pokemon{}
		UserPokemon.UserId = user.ID
		UserPokemon.PokemonId = starterPokemon[rand.IntN(3-0)+0]
		if DBClient.Create(&UserPokemon).Error != nil {
			log.Fatal(err)
			return
		}

		//TODO Update collection in Redis
	}
}
