package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func PrintAvailableCommands() {

	commands := LoadCommands()

	if len(commands) == 0 {
		Logger(LoggerOptions{Level: "warn", Message: "No commands was found in config"})
		return
	}

	headerColor := color.New(color.FgBlue, color.Bold)
	headerColor.Println("Available Commands:")

	fmt.Println(strings.Repeat("=", 40))

	for _, command := range commands {
		cmdNameColor := color.New(color.FgCyan, color.Bold)
		cmdNameColor.Printf("Command: %s\n", command.Name)

		fmt.Printf("String: %s\n", command.String)

		fmt.Println(strings.Repeat("-", 40))
	}

	Logger(LoggerOptions{Level: "info", Message: "(string means the command to be ran by the command runner)"})
}
