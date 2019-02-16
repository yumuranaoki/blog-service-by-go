package cacheserver

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

// RedisClient is client for redis
var RedisClient *redis.Client

// Init initialize cache server
func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	status, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("reids status: %s\n", status)
	defer RedisClient.Close()
}
