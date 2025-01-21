package utils

func RunAll() {
	for _, command := range ParseCommands() {
		RunCmd(command.String, true)
	}
}
