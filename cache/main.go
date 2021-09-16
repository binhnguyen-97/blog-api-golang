package cache

import (
	"blog-api-golang/config"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client
var ctx = context.Background()

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB,
	})
}

func Close() {
	if err := Client.Close(); err != nil {
		panic(err)
	}
}

func GetValue(key string) (string, error) {
	value, err := Client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}

func SetValue(key string, value interface{}, duration time.Duration) bool {

	log.Println("Set Value", key, value)
	err := Client.Set(ctx, key, value, duration).Err()

	if err != nil {
		log.Println("Error", err)
		return false
	}

	log.Println("Success", err)

	return true
}
