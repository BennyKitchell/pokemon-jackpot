package main

import (
	"log"
	"pokemon-service/controllers"
	initializers "pokemon-service/init"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	port      = ":8084"
	userTopic = "users"
)

func init() {
	//setup redis
	redisClient := initializers.ConnectRedis()
	controllers.SetRedis(redisClient)

	// setup kafka
	userTopicConsumer := initializers.ConnectConsumerToKafka()
	controllers.SetUserTopicConsumer(userTopicConsumer)

	//setup db
	dbClient := initializers.ConnectDB()
	controllers.SetDbClient(dbClient)

}
func main() {
	println("Server running on port", port)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/v1/pokemon/:id", controllers.GetPokemon)
	router.GET("/v1/collection/:userid", controllers.GetCollection)
	go controllers.StartUserCreationConsumer()
	if err := router.Run(port); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
