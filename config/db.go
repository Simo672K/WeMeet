package config

import (
	"WeMeet/pkg/utils"
	"strconv"

	"sync"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func DBConnect() *redis.Client {
	var once sync.Once
	envVars := utils.LoadEnvVars("REDIS_ADDRESS", "REDIS_PASSWORD", "REDIS_DB")
	redisDb, _ := strconv.Atoi(envVars["REDIS_DB"])

	once.Do(func() {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     envVars["REDIS_ADDRESS"],
			Password: envVars["REDIS_ADDRESS"],
			DB:       redisDb,
		})
	})

	return RedisClient
}
