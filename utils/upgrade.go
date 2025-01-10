package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Masterminds/semver"
)

func Upgrade() {

	currentVersion := LoadVersion()

	if currentVersion == " " {
		fmt.Println("Could not detect version")
		os.Exit(1)
	}

	// Get the latest github release and match it to the current version

	fmt.Println("Upgrading to latest version")

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

	fmt.Println("Latest version: " + release.TagName)
	fmt.Println("Current version: " + currentVersion)

	v, err := semver.NewVersion(release.TagName)

	if err != nil {
		panic(err)
	}

	isOutdated, err := semver.NewConstraint("> " + currentVersion)

	if err != nil {
		panic(err)
	}

	if isOutdated.Check(v) {
		fmt.Println("Downloading latest version")
		RunCmd("go install github.com/lassejlv/action@" + release.TagName)
		fmt.Println("Successfully updated to " + release.TagName)
	} else {
		fmt.Println("Your version is up to date")
	}

}
