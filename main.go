package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// fmt.Println(style.Render("Redis TUI"))
	name := scanText("Your name")
	fmt.Println("Hello", name)

	_, _, _ = connectToRedis()

}

func scanText(outputTxt string) string {
	var inputText string

	fmt.Println(outputTxt)
	fmt.Scanln(&inputText)

	return inputText
}

func connectToRedis() (*redis.Client, context.Context, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()

	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		return nil, nil, err
	} else {
		fmt.Println(pong)
	}

	return rdb, ctx, err
}
