package main

import (
	"fmt"
	"os"

	"github.com/lassejlv/action/utils"
)

func main() {

	commands := utils.LoadCommands(os.Args[1])

	// Run the first command if no command is specified
	if len(os.Args) < 2 {
		utils.RunCmd(commands[0].String)
		return
	}

	cmdToRun := os.Args[1]

	if cmdToRun == "--list" {
		utils.PrintAvailableCommands(cmdToRun)
		return
	}

	if cmdToRun == "--version" {
		utils.Logger(utils.LoggerOptions{
			Level:   "info",
			Message: fmt.Sprintf("Current version: %s", utils.CurrentVersion),
		})
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

	if cmdToRun == "--help" {
		utils.Usage()
		return
	}

	// Runs all the commands
	if cmdToRun == "--all" {

		for _, command := range commands {
			utils.RunCmd(command.String)
		}

		return
	}

	for _, command := range commands {
		if command.Name == cmdToRun {
			utils.RunCmd(command.String)
			return
		}
	}

	// fmt.Printf("Command '%s' not found in config\n", cmdToRun)
	utils.Logger(utils.LoggerOptions{
		Level:   "error",
		Message: fmt.Sprintf("Command '%s' not found in config", cmdToRun),
	})
}
