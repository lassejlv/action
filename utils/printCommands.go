package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func PrintAvailableCommands() {
	headerColor := color.New(color.FgBlue, color.Bold)
	headerColor.Println("Available Commands:")

	fmt.Println(strings.Repeat("=", 40))

	for _, command := range LoadCommands() {
		cmdNameColor := color.New(color.FgCyan, color.Bold)
		cmdNameColor.Printf("Command: %s\n", command.Name)

		fmt.Printf("String: %s\n", command.String)

		fmt.Println(strings.Repeat("-", 40))
	}
}
