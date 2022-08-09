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

type Veaver veave.Veaver

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

func EvalChecker(client *github.Client, veave *Veaver, ctx context.Context, owner, repo string) {
	for index, _ := range veave.Rules {
		mode := veave.Rules[index].Mode
		switch mode {
		case "backmerge":
			fmt.Println("mode: backmerge")
		case "sync":
			fmt.Println("mode: sync")
		case "in-sync":
			fmt.Println("mode: in-sync")
		case "in-sync-force":
			fmt.Println("mode: in-sync-force")
		case "removal":
			fmt.Println("mode: removal")
			// git.InSyncForce(client, ctx, owner, repo, source, destination_branches)

		}
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

func GetVeaverData(rawdata []byte) *Veaver {
	// Get Veaver Data Struct
	var data *Veaver
	data_err := yaml.Unmarshal(rawdata, &data)

	if data_err != nil {
		log.Fatal(data_err)
	}
	return data
}
