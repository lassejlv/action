package main

import (
	"os"

	"github.com/lassejlv/action/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Enable pretty logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	commands := utils.ParseCommands()

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
		log.Info().Msgf("v%s", utils.CurrentVersion)
		return
	}

	if cmdToRun == "--upgrade" {
		utils.Upgrade()
		return
	}

	if cmdToRun == "--init" {
		log.Info().Msg("Not implemented yet")
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

	log.Error().Msgf("Command '%s' not found in config", cmdToRun)
}
