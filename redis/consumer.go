package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func (r *RedisClient) Consume(ctx context.Context, id string) error {
	args := redis.XReadGroupArgs{
		Streams: []string{streamName, id},
		Group:   groupName,
		Block:   1 * time.Second,
		// Count: 1,
	}

	result, err := r.DB.XReadGroup(&args).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return err
		}
		return err
	}

	for _, s := range result {
		for _, message := range s.Messages {
			for _, v := range message.Values {
				var fruit Fruit
				err = json.Unmarshal([]byte(v.(string)), &fruit)
				if err != nil {
					continue
				}

				fmt.Printf("Fruit is consumed: %s\n", fruit.Name)
			}

			if err := r.DB.XAck(streamName, groupName, message.ID).Err(); err != nil {
				return fmt.Errorf("failed to acknowledge request with ID: %s, err: %v", message.ID, err)
			}
		}
	}
	return nil
}
