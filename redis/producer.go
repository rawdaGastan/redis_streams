package redis

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Fruit struct {
	Name string
}

func (r *RedisClient) Add(fruit Fruit) error {
	bytes, err := json.Marshal(fruit)
	if err != nil {
		return err
	}

	fmt.Printf("Fruit is added: %s\n", fruit.Name)

	return r.DB.XAdd(&redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{"fruit": bytes},
	}).Err()
}
