package utils

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

var ConfigFileName string = ".actions"
var CurrentVersion string = "1.0.3"

type CommandsArray struct {
	Name   string
	String string
}

// Example
// hello_world =  echo "Hello World"
// hello_mom = echo "Hello Mom, i love you!" {{ depends_on = "hello_world" }}
type CommandOptions struct {
	DependsOn string `json:"depends_on"`
}

func HasConfig(line string) (CommandOptions, error) {
	endsWith := strings.HasSuffix(line, "}")

	if !endsWith {
		return CommandOptions{}, nil
	}

	// Look for depends_on in the line
	if strings.Contains(line, "depends_on") {
		// Find the value between quotes after depends_on =
		startIndex := strings.Index(line, "depends_on") + len("depends_on")
		// Find the next quote after equals sign
		startIndex = strings.Index(line[startIndex:], "\"") + startIndex + 1
		if startIndex == -1 {
			return CommandOptions{}, nil
		}

		// Find the closing quote
		endIndex := strings.Index(line[startIndex:], "\"") + startIndex
		if endIndex == -1 {
			return CommandOptions{}, nil
		}

		dependsOn := line[startIndex:endIndex]
		return CommandOptions{
			DependsOn: dependsOn,
		}, nil
	}

	return CommandOptions{}, nil
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
				lineNumber = i + 1
			}
		}

		isEnv := strings.HasPrefix(line, "@env")

		if isEnv {
			// Remove the = and just get the path
			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				log.Error().Msgf("Invalid env format: %s at line %d", line, lineNumber)
				os.Exit(1)
			}
			envPath := strings.TrimSpace(parts[1])

			// Parse the env file
			content, err := os.ReadFile(envPath)
			if err != nil {
				log.Error().Msgf("Error reading env file: %s at line %d", line, lineNumber)
				os.Exit(1)
			}
			EnvParser(string(content))
			continue
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

		var finalCommand string

		isContainingEnv := strings.Contains(cmdString, "{") || strings.Contains(cmdString, "}")

		if isContainingEnv {
			// get the env betweeen {} and replace that
			envName := strings.Split(cmdString, "{")[1]
			envName = strings.Split(envName, "}")[0]
			envName = strings.TrimSpace(envName)
			envValue := os.Getenv(envName)
			finalCommand = strings.Replace(cmdString, "{"+envName+"}", envValue, 1)
		} else {
			finalCommand = cmdString
		}

		commands = append(commands, CommandsArray{Name: cmdName, String: finalCommand})
	}

	if len(commands) == 0 {
		log.Error().Msg("No valid commands found in config file")
		os.Exit(1)
	}

	return commands
}
