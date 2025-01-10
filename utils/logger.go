package utils

import (
	"fmt"

	"github.com/fatih/color"
)

type LoggerOptions struct {
	Level   string
	Message string
}

var (
	error   = color.New(color.FgRed, color.Bold)
	success = color.New(color.FgGreen, color.Bold)
	info    = color.New(color.FgBlue, color.Bold)
	warn    = color.New(color.FgYellow, color.Bold)
)

func Logger(options LoggerOptions) {
	switch options.Level {
	case "error":
		error.Println("[ERROR]", options.Message)
	case "success":
		success.Println(options.Message)
	case "info":
		info.Println("[INFO]", options.Message)
	case "warn":
		warn.Println("[WARN]", options.Message)
	default:
		fmt.Println(options.Message)
	}
}
