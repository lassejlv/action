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
  action <command>    Execute a specific command from .actions
  action --all       Run all available commands from .actions
  action --list      List all available commands from .actions
  action --version   Print the current version
  action --upgrade   Upgrade to the latest version
  action --init      Create a .actions file

Examples:
  action build       Run the build command - Loads the build command from .actions
  action test        Run the test command - Loads the test command from .actions

For more information, visit: https://github.com/lassejlv/actionfile
`

func Usage() {
	// Format usage with colors

	headerColor.Println("\nActionfile Command Runner")
	fmt.Printf(usageTemplate)
}
