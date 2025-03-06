package utils

func RunAll() {
	for _, command := range ParseCommands() {
		RunCmd(command.String, true)
	}

	for _, command := range ParsePackageJson() {
		RunCmd(command.String, true)
	}
}
