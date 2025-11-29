package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/redis/go-redis/v9"
)

func main() {

	_, _, _ = connectToRedis()

}

func scanText(outputTxt string) string {
	var inputText string

	fmt.Println(outputTxt)
	fmt.Scanln(&inputText)

	return inputText
}

func printWelcome(msg string) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF69B4")).
		Background(lipgloss.Color("#1E1E1E")).
		Padding(1, 4).
		Margin(1, 2).
		Border(lipgloss.NormalBorder(), true)

	fmt.Println(style.Render(msg))
}

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
