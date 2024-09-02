package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"pokemon-service/pkg/models"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
)

func GetPokemon(context *gin.Context) {
	pokemonId := context.Param("id")
	redisKey := fmt.Sprintf("pokemon-%s", pokemonId)

	// check cache
	val := RedisClient.Get(context, redisKey).Val()
	var pokemon models.Pokemon

	if len(val) != 0 {
		json.Unmarshal([]byte(val), &pokemon)
		context.JSON(http.StatusOK, pokemon)
		return
	}

	//fetch and hydrate cache
	if DBClient.Where("id = ?", pokemonId).First(&pokemon).Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	pokemonJSON, _ := json.Marshal(pokemon)
	_, err := RedisClient.Set(context, redisKey, pokemonJSON, 1*time.Minute).Result()
	if err != nil {
		log.Fatal(err)
	}
	if pokemon.ID == 0 {
		return
	}

	context.JSON(http.StatusOK, pokemon)
}

func RollPokemon(context *gin.Context) {
	spinNumber := context.Param("spinNumber")
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		log.Panic(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var randomPokemon []models.Pokemon
	pokemonIds := []string{}
	for range 3 {
		pokemonIds = append(pokemonIds, strconv.Itoa(rand.IntN(151)))
	}

	if DBClient.Where("id in (?)", pokemonIds).Find(&randomPokemon).Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Pokemon not found"})
		return
	}

	// If you rolled the same pokemon twice but not 3 times, handle not sending back 2 pokemon
	if len(randomPokemon) == 2 {
		randomPokemon[2] = randomPokemon[0]
	}

	if (randomPokemon[0].ID == randomPokemon[1].ID) && (randomPokemon[0].ID == randomPokemon[2].ID) {
		jackpot := models.User_Pokemon{UserId: user.ID, PokemonId: randomPokemon[0].ID}
		jackpotJson, err := json.Marshal(jackpot)
		if err != nil {
			log.Panic(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &JackpotTopic, Partition: kafka.PartitionAny},
			Key:            []byte([]byte("jackpot")),
			Value:          []byte(jackpotJson),
		}
		err = JackpotTopicProducer.Produce(message, nil)

		if err != nil {
			log.Printf("Error writing message to Kafka: %s", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"pokemon": randomPokemon, "jackpot": true})
		return
	}

	roll, err := strconv.Atoi(spinNumber)
	if err != nil {
		panic(err)
	}

	// If the roll count has been a few terms, give a jackpt for better user experience
	if roll%5 == 0 {
		for i := range 2 {
			randomPokemon[i] = randomPokemon[2]
		}

		tmp := models.User_Pokemon{UserId: user.ID, PokemonId: randomPokemon[0].ID}
		jackpotJson, err := json.Marshal(tmp)
		if err != nil {
			log.Panic(err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &JackpotTopic, Partition: kafka.PartitionAny},
			Key:            []byte([]byte("jackpot")),
			Value:          []byte(jackpotJson),
		}
		err = JackpotTopicProducer.Produce(message, nil)

		if err != nil {
			log.Printf("Error writing message to Kafka: %s", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"pokemon": randomPokemon, "jackpot": true})
		return
	}

	context.JSON(http.StatusOK, gin.H{"pokemon": randomPokemon})
}
