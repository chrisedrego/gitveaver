package utils

import (
	"fmt"
	"os"
)

var ConfigFile, GithubToken string

func HandleErr(err string) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func FlagCheck() error {
	// check ConfigFile, GithuToken
	ConfigFile = os.Getenv("CONFIGFILE")
	if ConfigFile == "" {
		ConfigFile = "/configs/veave.yaml"
	} else {
		fmt.Println("Default Configuration:", ConfigFile)
	}
	GithubToken = os.Getenv("GITHUBTOKEN")
	if GithubToken == "" {
		panic("No Value found for GITHUBTOKEN")
	}
	return nil
}

func CheckExecutor(branch, pushed_branch string) bool {
	// Check Executor
	fmt.Println("Veaver Branch: ", branch)
	fmt.Println("Pushed Branch: ", pushed_branch)
	if branch == pushed_branch {
		return true
	} else {
		return false
	}
}
