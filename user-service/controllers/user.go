package controllers

import (
	"log"
	"net/http"
	"user-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var user models.User
	if err := context.BindJSON(&user); err != nil {
		log.Panic(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if DBClient.Where("email = ?", user.Email).First(&user).Error == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists!"})
		return
	}

	dbError := DBClient.Create(&user).Error
	if dbError != nil {
		log.Fatal(dbError.Error())
	}

	// handle cache hydration here
	// handle kafka message here

	context.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
