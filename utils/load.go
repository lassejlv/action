package utils

import (
	"fmt"
	"os"
	"strings"
)

var ConfigFileName string = ".actions"
var CurrentVersion string = "0.1.15"

type CommandsArray struct {
	Name   string
	String string
}

func LoadCommands() []CommandsArray {

	fileExists, _ := os.Stat(ConfigFileName)

	if fileExists == nil {
		cwd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		fmt.Println("Could not detect .actions in ", cwd)
		os.Exit(1)
	}

	data, err := os.ReadFile(ConfigFileName)

	if err != nil {
		panic(err)
	}

	var commands []CommandsArray

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		cmdName := strings.TrimSpace(parts[0])
		cmdString := strings.TrimSpace(parts[1])

		commands = append(commands, CommandsArray{Name: cmdName, String: cmdString})
	}

	return commands
}

func LoadVersion() string {
	return CurrentVersion
}
