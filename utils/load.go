package utils

import (
	"os"
	"strings"
)

var ConfigFileName string = ".actions"

type CommandsArray struct {
	Name   string
	String string
}

func LoadCommands() []CommandsArray {
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
