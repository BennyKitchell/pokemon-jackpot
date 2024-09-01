package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"user-service/controllers"
	initializers "user-service/init"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const port = ":8020"

func init() {
	//setup db
	dbClient := initializers.ConnectDB()
	controllers.SetDbClient(dbClient)

	// setup redis
	redisClient := initializers.ConnectRedis()
	controllers.SetRedis(redisClient)

	//setup kafka
	userTopicProducer, err := initializers.ConnectProducerToKafka()
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	controllers.SetUserTopicProducer(userTopicProducer)
}

func main() {
	println("Server running on port: ", port)
	_, cancel := context.WithCancel(context.Background())

	defer cancel()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/v1/user", controllers.CreateUser)
	router.POST("/v1/login", controllers.GetUser)

	if err := router.Run(port); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
