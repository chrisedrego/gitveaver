package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/chrisedrego/gitveaver/cmd/gitveaver"
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
	var payload gitveaver.GithubPayload
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
	Owner, Repo := GetRepoDetails((string(payload.Repository.FullName)))

	// Retrieve Configuration Data
	var VeaverData *Veaver
	VeaverRawPayload := getRawVeaver(client, context, Owner, Repo, utils.ConfigFile, RefPushedBranch, RepGetOptions)
	fmt.Println(VeaverData, VeaverRawPayload)
	VeaverData.EvalChecker()
}
