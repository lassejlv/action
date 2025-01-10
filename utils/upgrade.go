package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Masterminds/semver"
)

func Upgrade() {

	currentVersion := LoadVersion()

	if currentVersion == " " {
		Logger(LoggerOptions{Level: "error", Message: "Could not detect version"})
		os.Exit(1)
	}

	// Get the latest github release and match it to the current version

	Logger(LoggerOptions{Level: "info", Message: "Checking for updates..."})

	url := "https://api.github.com/repos/lassejlv/actionfile-go/releases/latest"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	var release struct {
		TagName string `json:"tag_name"`
	}

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		panic(err)
	}

	Logger(LoggerOptions{Level: "info", Message: "Checking for updates..."})
	Logger(LoggerOptions{Level: "info", Message: "Latest version: " + release.TagName})
	Logger(LoggerOptions{Level: "info", Message: "Current version: " + currentVersion})

	v, err := semver.NewVersion(release.TagName)

	if err != nil {
		panic(err)
	}

	isOutdated, err := semver.NewConstraint("> " + currentVersion)

	if err != nil {
		panic(err)
	}

	if isOutdated.Check(v) {
		Logger(LoggerOptions{Level: "info", Message: "Upgrading to " + release.TagName})
		RunCmd("go install github.com/lassejlv/action@" + release.TagName)
		Logger(LoggerOptions{Level: "success", Message: "Upgrade complete"})
	} else {
		Logger(LoggerOptions{Level: "info", Message: "You are already on the latest version"})
	}

}
