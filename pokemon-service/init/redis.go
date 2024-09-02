package initializers

import (
	"context"

	redis_connection "pokemon-service/modules/redis"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	redisOptions := &redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	}

	return redis_connection.NewRedisClient(redisOptions, context.Background())
}
