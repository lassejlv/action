package utils

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

var ConfigFileName string = ".actions"
var CurrentVersion string = "0.1.28"

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
			log.Warn().Msgf("No %s file found in %s", ConfigFileName, cwd)
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
			log.Warn().Msgf("Invalid line format: %s", line)
			continue
		}

		cmdName := strings.TrimSpace(parts[0])
		cmdString := strings.TrimSpace(parts[1])

		// Remove comments from command string
		if commentIndex := strings.Index(cmdString, "#"); commentIndex != -1 {
			cmdString = strings.TrimSpace(cmdString[:commentIndex])
		}

		if cmdName == "" || cmdString == "" {
			log.Warn().Msgf("Empty command name or string: %s", line)
			continue
		}

		commands = append(commands, CommandsArray{Name: cmdName, String: cmdString})
	}

	if len(commands) == 0 {
		log.Warn().Msg("No valid commands found in config file")
	}

	return commands
}
