package utils

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/rs/zerolog/log"
)

func RunCmd(cmdString string) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdString)
	} else {
		cmd = exec.Command("sh", "-c", cmdString)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Error().Err(err).Msg("Could not run the command, sadly.")
		os.Exit(1)
	}
}
