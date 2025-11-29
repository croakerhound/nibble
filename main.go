package main

import (
	"context"
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/redis/go-redis/v9"
)

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	PaddingTop(2).
	PaddingLeft(4).
	Width(22)

func main() {
	// fmt.Println(style.Render("Redis TUI"))
	name := scanText("Your name")
	fmt.Println("Hello", name)

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password
		DB:       0,  // use default DB
		Protocol: 2,
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}

	fmt.Println("Connected to Redis:", pong)

}

func scanText(outputTxt string) string {
	var inputText string

	fmt.Println(outputTxt)
	fmt.Scanln(&inputText)

	return inputText
}
