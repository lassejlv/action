package main

import (
	"actionfile/utils"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

	// Runs all the commands
	if cmdToRun == "--all" {
		fmt.Println("Running all commands...")

		for _, command := range utils.LoadCommands() {
			var cmd *exec.Cmd

			if runtime.GOOS == "windows" {
				cmd = exec.Command("cmd", "/C", command.String)
			} else {
				cmd = exec.Command("sh", "-c", command.String)
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				os.Exit(1)
			}

		}
		return
	}

	for _, command := range utils.LoadCommands() {
		if command.Name == cmdToRun {
			var cmd *exec.Cmd

			if runtime.GOOS == "windows" {
				cmd = exec.Command("cmd", "/C", command.String)
			} else {
				cmd = exec.Command("sh", "-c", command.String)
			}

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				os.Exit(1)
			}
			return

		}
	}

	fmt.Printf("Command '%s' not found in config\n", cmdToRun)
}
