package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
