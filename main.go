package main

import (
	"fmt"
	"os"

	"github.com/lassejlv/action/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Enable pretty logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	commands := utils.ParseCommands()
	npm_scripts := utils.ParsePackageJson()

	// Run the first command if no command is specified
	if len(os.Args) < 2 && len(commands) > 0 {
		utils.RunCmd(commands[0].String, true)
		return
	}

	cmdToRun := os.Args[1]

	if cmdToRun == "--list" || cmdToRun == "ls" {
		utils.PrintAvailableCommands(cmdToRun)
		return
	}

	if cmdToRun == "--version" || cmdToRun == "-v" || cmdToRun == "version" {
		fmt.Println(utils.CurrentVersion)
		return
	}

	if cmdToRun == "--upgrade" || cmdToRun == "upgrade" || cmdToRun == "update" {
		utils.Upgrade()
		return
	}

	if cmdToRun == "--init " || cmdToRun == "init" {
		log.Info().Msg("Not implemented yet")
		return
	}

	if cmdToRun == "--help" || cmdToRun == "help" {
		utils.Help()
		return
	}

	if cmdToRun == "--all" || cmdToRun == "all" {
		utils.RunAll()
		return
	}

	for _, command := range commands {
		if command.Name == cmdToRun {
			config, err := utils.HasConfig(command.String)

			if err != nil {
				log.Error().Msgf("Error parsing config: %s", err)
				return
			} else {
				fmt.Print(config)
			}

			utils.RunCmd(command.String, true)
			return
		}
	}

	for _, script := range npm_scripts {
		if script.Name == cmdToRun {
			utils.RunCmd(script.String, true)
			return
		}
	}

	log.Error().Msgf("Command '%s' not found in config", cmdToRun)
}
