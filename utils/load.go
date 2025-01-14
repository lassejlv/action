package utils

import (
	"fmt"
	"os"
	"strings"
)

var ConfigFileName string = ".actions"
var CurrentVersion string = "0.1.24"

type CommandsArray struct {
	Name   string
	String string
}

func LoadCommands() []CommandsArray {
	// First check if file exists
	_, err := os.Stat(ConfigFileName)
	if err != nil {
		if os.IsNotExist(err) {
			cwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			Logger(LoggerOptions{Level: "warn", Message: fmt.Sprintf("No %s file found in %s", ConfigFileName, cwd)})
			os.Exit(1)
		}
		panic(err)
	}

	// Read file
	data, err := os.ReadFile(ConfigFileName)
	if err != nil {
		panic(err)
	}

	var commands []CommandsArray

	// Parse each line
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split by = first
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			Logger(LoggerOptions{Level: "warn", Message: fmt.Sprintf("Invalid line format: %s", line)})
			continue
		}

		cmdName := strings.TrimSpace(parts[0])
		cmdString := strings.TrimSpace(parts[1])

		// Remove comments from command string
		if commentIndex := strings.Index(cmdString, "#"); commentIndex != -1 {
			cmdString = strings.TrimSpace(cmdString[:commentIndex])
		}

		if cmdName == "" || cmdString == "" {
			Logger(LoggerOptions{Level: "warn", Message: fmt.Sprintf("Empty command name or string: %s", line)})
			continue
		}

		commands = append(commands, CommandsArray{Name: cmdName, String: cmdString})
	}

	if len(commands) == 0 {
		Logger(LoggerOptions{Level: "warn", Message: "No valid commands found in config file"})
	}

	return commands
}

func LoadVersion() string {
	return CurrentVersion
}
