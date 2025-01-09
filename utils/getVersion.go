package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var GithubRepoUrl string = "https://github.com/lassejlv/actionfile-go"

func GetLatestReleaseVersion() string {
	// Extract owner and repo from the URL
	parts := strings.Split(strings.TrimPrefix(GithubRepoUrl, "https://github.com/"), "/")
	if len(parts) != 2 {
		return ""
	}
	owner, repo := parts[0], parts[1]

	// Create GitHub API URL for latest release
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	// Make HTTP request
	resp, err := http.Get(apiUrl)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// Parse JSON response
	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.Unmarshal(body, &release); err != nil {
		return ""
	}

	// Return the version number without 'v' prefix if present
	return "v" + strings.TrimPrefix(release.TagName, "v")
}
