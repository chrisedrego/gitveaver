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

func RequestHandler(resp http.ResponseWriter, req *http.Request) {
	// handles request for GitVeaver

	// extract Github payload for events from webhook
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var payload veave.GithubPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		panic(err)
	}
	// extract branch details
	RefPushedBranch := (string(payload.Ref))
	// PushedBranch := GetBranch(string(payload.Ref))

	context := context.Background()

	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: utils.GithubToken},
	)
	tokenClient := oauth2.NewClient(context, tokenService)

	// Intiating Github Client
	client := github.NewClient(tokenClient)

	RepGetOptions := github.RepositoryContentGetOptions{
		Ref: RefPushedBranch,
	}

	// Getting Repo & Owner details
	Owner, Repo := git.GetRepoDetails((string(payload.Repository.FullName)))

	// Retrieve Configuration Data
	var VeaverData *utils.Veaver
	VeaverRawPayload := utils.GetRawVeaver(client, context, Owner, Repo, utils.ConfigFile, RefPushedBranch, RepGetOptions)
	VeaverData = (*utils.Veaver)(utils.GetVeaverData(VeaverRawPayload))
	fmt.Println(VeaverData, VeaverRawPayload)
	utils.EvalChecker(client, VeaverData, context, Owner, Repo)
}
