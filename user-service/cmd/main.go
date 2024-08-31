package main

import (
	"context"
	"log"
	"pokemon-jackpot/controllers"

	"github.com/gin-gonic/gin"
)

const port = ":8020"

func main() {
	println("Server running on port: ", port)
	_, cancel := context.WithCancel(context.Background())

	defer cancel()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/user", controllers.CreateUser)

	if err := router.Run(port); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
