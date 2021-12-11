package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func Conn() {
	redisDB, _ := strconv.Atoi(os.Getenv("redis_database"))
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_host") + ":" + os.Getenv("redis_port"),
		Password: os.Getenv("redis_password"),
		DB:       redisDB,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("error initializing Redis Client: %v\n", err)
	}
	fmt.Println("⇨ Redis Client Started !!")
}
