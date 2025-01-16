package utils

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
)

func PrintAvailableCommands(cmdToRun string) {

	commands := ParseCommands()

	if len(commands) == 0 {
		log.Warn().Msg("No commands was found in config")
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
}
