package utils

import (
	"encoding/json"
	"os"
)

// A function that checks if there is an package.json file and it contains "scripts". Then it will parse it and return them as an json obj
type Script struct {
	Name   string `json:"name"`
	String string `json:"string"`
}

func ParsePackageJson() []Script {
	// First check if file exists
	_, err := os.Stat("package.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		panic(err)
	}

	// Read file
	data, err := os.ReadFile("package.json")
	if err != nil {
		panic(err)
	}

	// Parse file
	var packageJSON struct {
		Scripts map[string]string `json:"scripts"`
	}
	
	err = json.Unmarshal(data, &packageJSON)
	if err != nil {
		panic(err)
	}

	var scriptsArray []Script
	for name, script := range packageJSON.Scripts {
		scriptsArray = append(scriptsArray, Script{
			Name:   name,
			String: script,
		})
	}

	return scriptsArray
}
