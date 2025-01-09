package utils

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

var errorColor = color.New(color.FgRed, color.Bold)

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
		errorColor.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
