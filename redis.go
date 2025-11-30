package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func ConnectToRedis() (*redis.Client, context.Context, error) {
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
