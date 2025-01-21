package utils

import (
	"os"
	"strings"
)

// A function that will parse the environment variables from the content and set them
func EnvParser(fileData string) {

	for _, line := range strings.Split(fileData, "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		os.Setenv(parts[0], parts[1])
	}

}
