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
	error   = color.New(color.FgRed)
	success = color.New(color.FgGreen)
	info    = color.New(color.FgBlue)
	warn    = color.New(color.FgYellow)
)

func Logger(options LoggerOptions) {
	switch options.Level {
	case "error":
		error.Println("[ERROR]", options.Message)
	case "success":
		success.Println("[SUCCESS]", options.Message)
	case "info":
		info.Println("[INFO]", options.Message)
	case "warn":
		warn.Println("[WARN]", options.Message)
	default:
		fmt.Println(options.Message)
	}
}
