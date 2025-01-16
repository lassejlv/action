package utils

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
)

func RunAll() {

	commands := ParseCommands()
	startTime := time.Now()
	log.Info().Msg("Running all commands")

	for _, command := range commands {
		RunCmd(command.String, true)
	}

	elapsed := time.Since(startTime).Seconds()
	log.Info().Msgf("Took %f seconds to run all commands", elapsed)
	log.Info().Msg("All commands ran successfully, with a total of " + fmt.Sprint(len(commands)) + " commands")
}
