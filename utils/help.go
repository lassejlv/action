package utils

import (
	"fmt"

	"github.com/fatih/color"
)

var headerColor = color.New(color.FgCyan, color.Bold)
var versionColor = color.New(color.FgGreen, color.Bold)

const usageTemplate = `
Usage:
  action <command>   Execute a specific command from .actions
  action --all       Run all available commands from .actions
  action --list      List all available commands from .actions
  action --version   Print the current version
  action --upgrade   Upgrade to the latest version
  action --help      Show this help message

For more information, visit: https://github.com/lassejlv/action
`

func Help() {
	headerColor.Println("\nActionfile Command Runner")
	versionColor.Printf("Version: %s\n", CurrentVersion)

	fmt.Print(usageTemplate)
}
