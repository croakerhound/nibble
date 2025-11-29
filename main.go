package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/redis/go-redis/v9"
)

// Context is needed for Redis operations (handles timeouts/cancellations)
var ctx = context.Background()

// Define styles using Lipgloss (like CSS classes)
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205")).
			MarginBottom(1)

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("86"))

	resultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(1).
			MarginBottom(1)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)
)

// model holds the state of our TUI application
// In Bubble Tea, this is like your React component's state
type model struct {
	redisClient *redis.Client   // Redis connection
	textInput   textinput.Model  // Text input field
	result      string           // Command results to display
	err         error            // Error messages
	connected   bool             // Connection status
}

// initialModel creates the starting state of our app
func initialModel() model {
	// Create text input component
	ti := textinput.New()
	ti.Placeholder = "Enter Redis command (e.g., SET key value, GET key)"
	ti.Focus() // Make it active immediately
	ti.CharLimit = 256
	ti.Width = 50

	// Configure Redis client
	// Change these values to match your Redis setup
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password by default
		DB:       0,                // Use default DB
	})

	// Test connection
	_, err := rdb.Ping(ctx).Result()
	connected := err == nil

	return model{
		redisClient: rdb,
		textInput:   ti,
		result:      "Connected to Redis! Type a command and press Enter.",
		connected:   connected,
		err:         err,
	}
}

// Init is called when the program starts
// It can return a command to run (we don't need any)
func (m model) Init() tea.Cmd {
	return textinput.Blink // Make cursor blink
}

// Update handles all user input and events
// This is like your event handlers in other frameworks
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// Handle keyboard input
	case tea.KeyMsg:
		switch msg.Type {
		// Quit the program
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		// Execute Redis command
		case tea.KeyEnter:
			if !m.connected {
				m.err = fmt.Errorf("not connected to Redis")
				return m, nil
			}

			// Get the command from input
			command := m.textInput.Value()
			if command == "" {
				return m, nil
			}

			// Parse and execute command
			m.result, m.err = m.executeRedisCommand(command)

			// Clear input for next command
			m.textInput.Reset()
		}
	}

	// Update the text input component
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// executeRedisCommand parses and runs a Redis command
func (m model) executeRedisCommand(command string) (string, error) {
	// Split command into parts (e.g., "SET key value" -> ["SET", "key", "value"])
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "", fmt.Errorf("empty command")
	}

	// Convert to uppercase for consistency
	cmd := strings.ToUpper(parts[0])
	args := parts[1:]

	// Handle different Redis commands
	switch cmd {
	case "GET":
		if len(args) != 1 {
			return "", fmt.Errorf("GET requires exactly 1 argument")
		}
		val, err := m.redisClient.Get(ctx, args[0]).Result()
		if err == redis.Nil {
			return "(nil)", nil
		}
		return val, err

	case "SET":
		if len(args) < 2 {
			return "", fmt.Errorf("SET requires at least 2 arguments")
		}
		// Join remaining args as value (in case value has spaces)
		value := strings.Join(args[1:], " ")
		return m.redisClient.Set(ctx, args[0], value, 0).Result()

	case "DEL":
		if len(args) == 0 {
			return "", fmt.Errorf("DEL requires at least 1 argument")
		}
		deleted, err := m.redisClient.Del(ctx, args...).Result()
		return fmt.Sprintf("Deleted %d key(s)", deleted), err

	case "KEYS":
		if len(args) != 1 {
			return "", fmt.Errorf("KEYS requires exactly 1 argument")
		}
		keys, err := m.redisClient.Keys(ctx, args[0]).Result()
		if err != nil {
			return "", err
		}
		return strings.Join(keys, "\n"), nil

	case "PING":
		return m.redisClient.Ping(ctx).Result()

	case "DBSIZE":
		size, err := m.redisClient.DBSize(ctx).Result()
		return fmt.Sprintf("%d keys", size), err

	case "FLUSHDB":
		return m.redisClient.FlushDB(ctx).Result()

	default:
		// For other commands, use generic Do method
		// Convert args to interface{} slice
		genericArgs := make([]interface{}, len(parts))
		for i, arg := range parts {
			genericArgs[i] = arg
		}
		result, err := m.redisClient.Do(ctx, genericArgs...).Result()
		return fmt.Sprintf("%v", result), err
	}
}

// View renders the UI
// This is like your render function - it returns a string that gets displayed
func (m model) View() string {
	// Build the UI string
	var s strings.Builder

	// Title
	s.WriteString(titleStyle.Render("🔴 Redis TUI"))
	s.WriteString("\n\n")

	// Connection status
	if m.connected {
		s.WriteString(inputStyle.Render("● Connected"))
	} else {
		s.WriteString(errorStyle.Render("● Disconnected"))
	}
	s.WriteString("\n\n")

	// Input field
	s.WriteString(inputStyle.Render("Command:") + "\n")
	s.WriteString(m.textInput.View())
	s.WriteString("\n\n")

	// Results or errors
	if m.err != nil {
		s.WriteString(errorStyle.Render("Error: " + m.err.Error()))
	} else if m.result != "" {
		s.WriteString(resultStyle.Render("Result:\n" + m.result))
	}

	s.WriteString("\n\n")

	// Help text
	s.WriteString(helpStyle.Render("Enter: Execute | Esc/Ctrl+C: Quit"))
	s.WriteString("\n")
	s.WriteString(helpStyle.Render("Examples: GET mykey | SET mykey myvalue | KEYS * | PING"))

	return s.String()
}

func main() {
	// Create the program
	p := tea.NewProgram(initialModel())

	// Run it (blocks until program exits)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
