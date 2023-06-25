package main

import (
	"context"
	"fmt"

	"github.com/rawdaGastan/redis_streams/redis"
)

func main() {
	// Connect to redis
	redisDB, err := redis.NewRedisClient("localhost", "6379", "")
	if err != nil {
		fmt.Println(err)
	}

	for {
		var fruit string
		fmt.Println("Please enter the name of the fruit")
		_, err = fmt.Scanln(&fruit)
		if err != nil {
			fmt.Println(err)
		}

		err = redisDB.Add(redis.Fruit{Name: fruit})
		if err != nil {
			fmt.Println(err)
		}

		//  change ">" to "0" to consume pending requests
		redisDB.Consume(context.Background(), ">")
		fmt.Println("-------------------------------")
	}
}
