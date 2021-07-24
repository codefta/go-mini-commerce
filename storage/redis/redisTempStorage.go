package redis_storage

import (
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis"
)

type RedisStorage struct {
	client *redis.Client
}

type RedisStorageConfigs struct {
	RedisHost	string
	RedisPort	string
	RedisPass	string
}

func NewRedisStorage(configs RedisStorageConfigs) (*RedisStorage, error) {
	client := redis.NewClient(&redis.Options{
		Addr: configs.RedisHost + ":" + configs.RedisPort,
		Password: configs.RedisPass,
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("unable to ping redis due to: %v", err)
	}

	return &RedisStorage{client: client}, nil
}