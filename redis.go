package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func connectToRedis() (*redis.Client, context.Context, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		printWelcome("Failed to connect to redis")
		return nil, nil, err
	} else {
		printWelcome("Connected to redis")
	}

	return rdb, ctx, err
}

func getAllKeys(rdb *redis.Client, ctx context.Context) {
	keys, err := rdb.Keys(ctx, "*").Result()

	if err != nil {
		fmt.Println("Failed to get Keys", err)
	}

	fmt.Println("Keys in Redis")
	for _, key := range keys {
		fmt.Println("-", key)
	}
}
