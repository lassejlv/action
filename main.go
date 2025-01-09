package main

import (
	"actionfile/utils"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command, use --list to list commands")
		os.Exit(1)
	}

	cmdToRun := os.Args[1]

	if cmdToRun == "--list" {
		utils.PrintAvailableCommands()
		return
	}

	if cmdToRun == "--version" {
		fmt.Println(utils.GetLatestReleaseVersion())
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
