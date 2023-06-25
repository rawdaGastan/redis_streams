package redis

import (
	"github.com/go-redis/redis"
)

const (
	streamName = "mystream"
	groupName  = "mygroup"
)

// RedisClient for redis DB handling streams
type RedisClient struct {
	DB *redis.Client
}

// NewRedisClient creates a new RedisClient
func NewRedisClient(redisHost, redisPort, redisPass string) (RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return RedisClient{}, err
	}

	client.XGroupCreateMkStream(streamName, groupName, "$")

	return RedisClient{client}, nil
}
