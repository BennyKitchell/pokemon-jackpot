package main

import (
	"context"
	"log"
	"user-service/controllers"
	initializers "user-service/init"

	"github.com/gin-gonic/gin"
)

const port = ":8020"

func init() {
	//setup db
	dbClient := initializers.ConnectDB()
	controllers.SetDbClient(dbClient)

	// setup redis

	//setup kafka
}

func main() {
	println("Server running on port: ", port)
	_, cancel := context.WithCancel(context.Background())

	defer cancel()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/v1/user", controllers.CreateUser)

	if err := router.Run(port); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
