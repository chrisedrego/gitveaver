package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chrisedrego/gitveaver/pkg/veave"
	"github.com/google/go-github/github"
	"gopkg.in/yaml.v2"
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

func GetRawVeaver(client *github.Client, ctx context.Context, owner, repo, filepath, branch string, RepGetOptions github.RepositoryContentGetOptions) []byte {
	// Get Contents of GitVeaver Configuration
	FileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filepath, &RepGetOptions)
	if err != nil {
		HandleErr(err.Error())
	}
	rawDecodedData, err := base64.StdEncoding.DecodeString(*FileContent.Content)
	if err != nil {
		panic(err)
	}
	return rawDecodedData
}

func CheckVeavied(data string) bool {
	var GV_PRFLag string
	return strings.HasPrefix(data, GV_PRFLag)
}

func GetVeaverData(rawdata []byte) *veave.Veaver {
	// Get Veaver Data Struct
	var data *veave.Veaver
	data_err := yaml.Unmarshal(rawdata, &data)

	if data_err != nil {
		log.Fatal(data_err)
	}
	return data
}
