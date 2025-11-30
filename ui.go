package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

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

func printKeys(keys []string) {
	if len(keys) == 0 {
		fmt.Println("No keys found.")
		return
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Padding(0, 1)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Padding(1, 2).
		Width(30)

	// Key base style (common)
	baseStyle := lipgloss.NewStyle().
		PaddingLeft(2).
		Width(20)

	// Two alternating row styles (zebra stripes)
	evenRowStyle := baseStyle.
		Foreground(lipgloss.Color("#00E676")). // bright green
		Background(lipgloss.Color("#1E1E1E"))  // dark gray background

	oddRowStyle := baseStyle.
		Foreground(lipgloss.Color("#80CBC4")). // teal
		Background(lipgloss.Color("#121212"))  // slightly darker bg

	// Build striped list
	var styledKeys []string
	for i, key := range keys {
		entry := fmt.Sprintf("%2d. %s", i+1, key)

		if i%2 == 0 {
			styledKeys = append(styledKeys, evenRowStyle.Render(entry))
		} else {
			styledKeys = append(styledKeys, oddRowStyle.Render(entry))
		}
	}

	keysBlock := strings.Join(styledKeys, "\n")

	final := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("🔑 Redis Keys"),
		borderStyle.Render(keysBlock),
	)

	fmt.Println(final)
}
