package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chrisedrego/gitveaver/pkg/git"
	"github.com/chrisedrego/gitveaver/pkg/veave"
	"github.com/chrisedrego/gitveaver/utils"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func getGitClient(context context.Context) *github.Client {

	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: utils.GithubToken},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	client := github.NewClient(tokenClient)
	return client
}

func RequestHandler(resp http.ResponseWriter, req *http.Request) {
	// handles request for GitVeaver

	// extract Github payload for events from webhook
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var payload *veave.GithubPayload

	err = json.Unmarshal(body, &payload)
	if err != nil {
		fmt.Println("Unable to Marshal JSON Body", err)
	}

	// extract branch details
	RefPushedBranch := (string(payload.Ref))

	PushedBranch := git.GetBranch(string(payload.Ref))
	fmt.Println(RefPushedBranch, PushedBranch)

	context := context.Background()

	RepGetOptions := github.RepositoryContentGetOptions{
		Ref: RefPushedBranch,
	}
	client := getGitClient(context)

	// Getting Repo & Owner details
	fmt.Println("Raw Payload", (string(payload.Repository.Full_Name)))
	Owner, Repo := git.GetRepoDetails((string(payload.Repository.Full_Name)))

	// Retrieve Configuration Data
	var VeaverData *veave.Veaver
	VeaverRawPayload := utils.GetRawVeaver(client, context, Owner, Repo, utils.ConfigFile, RefPushedBranch, RepGetOptions)
	VeaverData = (*veave.Veaver)(utils.GetVeaverData(VeaverRawPayload))
	EvalChecker(client, VeaverData, context, Owner, Repo)
}

func EvalChecker(client *github.Client, veave *veave.Veaver, ctx context.Context, owner, repo string) {
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
			git.InSyncForce(client, ctx, owner, repo, veave.Rules[index].SourceBranch, veave.Rules[index].DestinationBranches, veave.Rules[index].BranchProtection)
		case "removal":
			fmt.Println("mode: removal")
			git.Remove(client, ctx, owner, repo, veave, github.RepositoryContentGetOptions{})
		}
	}
}
