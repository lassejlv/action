package utils

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

var ConfigFileName string = ".actions"
var CurrentVersion string = "1.0.1"

type CommandsArray struct {
	Name   string
	String string
}

func ParseCommands() []CommandsArray {
	// First check if file exists
	_, err := os.Stat(ConfigFileName)
	if err != nil {
		if os.IsNotExist(err) {
			cwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			log.Error().Msgf("No %s file found in %s", ConfigFileName, cwd)
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

		var lineNumber int

		// get the line number
		for i, v := range strings.Split(string(data), "\n") {
			if v == line {
				lineNumber = i
			}
		}

		// Split by = first
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			log.Error().Msgf("Invalid line format: %s at line %d", line, lineNumber)
			os.Exit(1)
			continue
		}

		cmdName := strings.TrimSpace(parts[0])
		cmdString := strings.TrimSpace(parts[1])

		// Remove comments from command string
		if commentIndex := strings.Index(cmdString, "#"); commentIndex != -1 {
			cmdString = strings.TrimSpace(cmdString[:commentIndex])
		}

		if cmdName == "" || cmdString == "" {
			log.Error().Msgf("Empty command name or string: %s at line %d", line, lineNumber)
			os.Exit(1)
			continue
		}

		commands = append(commands, CommandsArray{Name: cmdName, String: cmdString})
	}

	if len(commands) == 0 {
		log.Error().Msg("No valid commands found in config file")
		os.Exit(1)
	}

	return commands
}
