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

	if len(commands) == 0 {
		log.Warn().Msg("No commands was found in config")
		return
	}

	// Create header
	fmt.Println("Available Commands:")

	// Create tabwriter for aligned output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Print each command with description
	for _, command := range commands {
		cmdColor := color.New(color.FgCyan, color.Bold)
		fmt.Fprintf(w, "  %s\t%s\n",
			cmdColor.Sprint(command.Name),
			command.String,
		)
	}
	w.Flush()
}
