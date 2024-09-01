package controllers

import (
	"fmt"
	"net/http"
	"pokemon-service/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCollection(context *gin.Context) {
	userId := context.Param("userid")
	var collection []models.User_Pokemon
	var ownedPokemon []models.Pokemon

	if DBClient.Where("user_id = ?", userId).Find(&collection).Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Collection not found"})
		return
	}

	fmt.Printf("collection is %d\n", len(collection))

	mapped := []string{}
	for i := range len(collection) {
		mapped = append(mapped, strconv.Itoa(collection[i].PokemonId))
	}

	if DBClient.Where("id in (?)", mapped).Find(&ownedPokemon).Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Owned Pokemon not found"})
		return
	}

	//TODO: hydrate cache

	context.JSON(http.StatusOK, gin.H{"pokemon": ownedPokemon})
}
