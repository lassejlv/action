package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	//	errorColor   = color.New(color.FgRed, color.Bold)
	headerColor  = color.New(color.FgCyan, color.Bold)
	commandColor = color.New(color.FgYellow)
	exampleColor = color.New(color.FgBlack)
	cliName      = "action"
)

const usageTemplate = `
Usage:
  %s <command>    Execute a specific command from .actions
  %s --all       Run all available commands from .actions
	%s --list      List all available commands from .actions
	%s --version   Print the current version
	%s --upgrade   Upgrade to the latest version
	%s --init      Create a .actions file

Examples:
  %s build       Run the build command - Loads the build command from .actions
  %s test        Run the test command - Loads the test command from .actions

For more information, visit: https://github.com/lassejlv/actionfile
`

func Usage() {
	// Format usage with colors
	usage := fmt.Sprintf(usageTemplate,
		commandColor.Sprint(cliName),
		commandColor.Sprint(cliName),
		exampleColor.Sprint(cliName),
		exampleColor.Sprint(cliName),
		exampleColor.Sprint(cliName),
		exampleColor.Sprint(cliName),
		exampleColor.Sprint(cliName),
		commandColor.Sprint(cliName),
	)

	headerColor.Println("\nActionfile Command Runner")
	fmt.Println(usage)
}
