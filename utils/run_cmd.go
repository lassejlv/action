package utils

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/rs/zerolog/log"
)

func RunCmd(cmdString string, showOutput bool) {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", cmdString)
	} else {
		cmd = exec.Command("sh", "-c", cmdString)
	}

	if showOutput {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	err := cmd.Run()

	if err != nil {
		log.Error().Msgf("Error running command: %s", err)
		os.Exit(1)
	}
}
