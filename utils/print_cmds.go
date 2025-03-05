package utils

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
)

func PrintAvailableCommands(cmdToRun string) {
	commands := ParseCommands()
	npm_scripts := ParsePackageJson()

	if len(commands) == 0 && len(npm_scripts) == 0 {
		log.Warn().Msg("No commands found in config or package.json")
		return
	}

	// Convert npm_scripts to CommandsArray type
	scriptCommands := make([]CommandsArray, len(npm_scripts))
	for i, script := range npm_scripts {
		scriptCommands[i] = CommandsArray{
			Name:   script.Name,
			String: script.String,
		}
	}

	// Combine both slices
	all_commands := append(commands, scriptCommands...)

	// Create header
	fmt.Println("Available Commands:")

	// Create tabwriter for aligned output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Print each command with description
	for _, command := range all_commands {
		cmdColor := color.New(color.FgCyan, color.Bold)
		fmt.Fprintf(w, "  %s\t%s\n",
			cmdColor.Sprint(command.Name),
			command.String,
		)
	}
	w.Flush()
}
