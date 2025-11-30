package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func main() {

	_, _, _ = ConnectToRedis()

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
