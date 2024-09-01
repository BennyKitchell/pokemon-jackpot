package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pokemon-service/pkg/models"
	"time"

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
