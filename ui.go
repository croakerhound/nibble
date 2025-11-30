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

func printKeysAndValues(data map[string]string) {
	if len(data) == 0 {
		fmt.Println("No keys or values found.")
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
		Width(60)

	// Base row style
	baseStyle := lipgloss.NewStyle().
		PaddingLeft(2).
		Width(55)

	evenRowStyle := baseStyle.
		Foreground(lipgloss.Color("#00E676")). // bright green
		Background(lipgloss.Color("#1E1E1E"))

	oddRowStyle := baseStyle.
		Foreground(lipgloss.Color("#80CBC4")). // teal
		Background(lipgloss.Color("#121212"))

	// Build zebra-striped key-value list
	var styledRows []string
	i := 0
	for key, value := range data {
		entry := fmt.Sprintf("%2d. %s → %s", i+1, key, value)

		if i%2 == 0 {
			styledRows = append(styledRows, evenRowStyle.Render(entry))
		} else {
			styledRows = append(styledRows, oddRowStyle.Render(entry))
		}
		i++
	}

	list := strings.Join(styledRows, "\n")

	final := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("🧠 Redis Keys & Values"),
		borderStyle.Render(list),
	)

	fmt.Println(final)
}
