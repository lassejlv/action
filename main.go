package main

import (
	"fmt"
	"os"

	"github.com/lassejlv/action/utils"

	"github.com/fatih/color"
)

var (
	errorColor   = color.New(color.FgRed, color.Bold)
	headerColor  = color.New(color.FgCyan, color.Bold)
	commandColor = color.New(color.FgYellow)
	exampleColor = color.New(color.FgBlack)
)

const usageTemplate = `
Usage:
  %s <command>    Execute a specific command from .actions
  %s --all       Run all available commands from .actions

Examples:
  %s build       Run the build command - Loads the build command from .actions
  %s test        Run the test command - Loads the test command from .actions

For more information, visit: https://github.com/lassejlv/actionfile-go
`

func main() {

	if len(os.Args) < 2 {
		errorColor.Fprintln(os.Stderr, "Error: No command specified")

		// Format usage with colors
		usage := fmt.Sprintf(usageTemplate,
			commandColor.Sprint("action"),
			commandColor.Sprint("action"),
			exampleColor.Sprint("action"),
			exampleColor.Sprint("action"),
		)

		headerColor.Println("\nActionfile Task Runner")
		fmt.Println(usage)
		os.Exit(1)
	}
	cmdToRun := os.Args[1]

	if cmdToRun == "--list" {
		utils.PrintAvailableCommands()
		return
	}

	if cmdToRun == "--version" {
		fmt.Println(utils.LoadVersion())
		return
	}

	if cmdToRun == "--upgrade" {
		utils.Upgrade()
		return
	}

	if cmdToRun == "--init" {
		utils.Init()
		return
	}

	// Runs all the commands
	if cmdToRun == "--all" {

		for _, command := range utils.LoadCommands() {
			utils.RunCmd(command.String)
		}

		return
	}

	for _, command := range utils.LoadCommands() {
		if command.Name == cmdToRun {
			utils.RunCmd(command.String)
			return
		}
	}

	fmt.Printf("Command '%s' not found in config\n", cmdToRun)
}
