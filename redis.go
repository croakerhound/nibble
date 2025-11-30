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

func getAllKeys(rdb *redis.Client, ctx context.Context) []string {
	keys, err := rdb.Keys(ctx, "*").Result()

	if err != nil {
		fmt.Println("Failed to get Keys", err)
	}
	return keys
}

func getAllKeysAndValues(rdb *redis.Client, ctx context.Context) (map[string]string, error) {
	// Get all keys
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}

	// Prepare a map to store key-value pairs
	data := make(map[string]string)

	// Loop through each key and fetch its value
	for _, key := range keys {
		val, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			// Key doesn't have a value (e.g. was deleted)
			continue
		} else if err != nil {
			return nil, fmt.Errorf("failed to get value for key %s: %w", key, err)
		}
		data[key] = val
	}

	return data, nil
}
