package utils

import (
	"os"
)

var ConfigFile, GithubToken string

func FlagCheck() error {
	// check ConfigFile, GithuToken
	ConfigFile = os.Getenv("CONFIGFILE")
	if ConfigFile == "" {
		ConfigFile = "config.yaml"
	}
	GithubToken = os.Getenv("GITHUBTOKEN")
	if GithubToken == "" {
		panic("No Value found for GITHUBTOKEN")
	}
	return nil
}
