package redis

import (
	"context"
	"log"

	"github.com/bsm/redislock"
	"github.com/cherifBurette1/rails-go-chat/tree/master/go-chat-creation-api/configs"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var redisLocker *redislock.Client
var ctx = context.Background()

func GetRedis() (*redis.Client, error) {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     configs.RedisAddress,
			Password: "",
			DB:       0,
		})

		err := redisClient.Ping(ctx).Err()
		if err != nil {
			return nil, err
		}
		redisLocker = redislock.New(redisClient)
	}
	return redisClient, nil
}

func GetLocker() *redislock.Client {
	if redisClient == nil {
		log.Fatalln("Redis client is not initialized yet")
	}
	return redisLocker
}
