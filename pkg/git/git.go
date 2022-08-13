package git

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chrisedrego/gitveaver/pkg/veave"
	"github.com/chrisedrego/gitveaver/utils"
	"github.com/google/go-github/github"
)

func InSync() {
	// InSync between same
	// git clone
	// git merge source to destination
	// git commit && git push
}

func DeleteBranch(client *github.Client, ctx context.Context, owner, repo, branch string) (*github.Response, error) {
	// Delete Branch
	u := fmt.Sprintf("repos/%v/%v/git/refs/heads/%v", owner, repo, branch)
	req, err := client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	// // TODO: remove custom Accept header when this API fully launches
	// req.Header.Set("Accept", mediaTypeRequiredApprovingReviewsPreview)

	return client.Do(ctx, req, nil)
}

func CreateBranch(client *github.Client, ctx context.Context, owner, repo, source_branch, destination_branch string) {
	// (*github.Response, error)
	// Get SHA from source branch
	service := client.Git
	fmt.Println("Service:", service, "ctx", ctx, "Owner", owner, "Repo", repo)
	ref, resp, error := service.GetRef(ctx, owner, repo, "heads/"+source_branch)
	SHA := ref.Object.GetSHA()

	if error != nil {
		fmt.Println("Status Code:", resp.StatusCode, "error:", error)
	}
	fmt.Println("SHA:", SHA, "Ref:", ref.GetRef())
	*ref.Ref = "refs/heads/" + destination_branch
	fmt.Println("Ref", *ref.Ref)
	ref, resp, err := service.CreateRef(ctx, owner, repo, ref)
	if err != nil {
		fmt.Println("Status Code:", resp.StatusCode, "error:", error)
	}

	// u := fmt.Sprintf("repos/%v/%v/git/refs/heads/%v", owner, repo, source)
	// req, err := client.NewRequest("POST", u, nil)
	// if err != nil {
	// 	return nil, err
	// }

	// TODO: remove custom Accept header when this API fully launches
	// req.Header.Set("Accept", mediaTypeRequiredApprovingReviewsPreview)

	// return client.Do(ctx, req, nil)
}

func RemoveBranchProtection(client *github.Client, ctx context.Context, owner, repo, branch string) (*github.Response, error) {
	u := fmt.Sprintf("repos/%v/%v/git/refs/heads/%v", owner, repo, branch)
	req, err := client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	// // TODO: remove custom Accept header when this API fully launches
	// req.Header.Set("Accept", mediaTypeRequiredApprovingReviewsPreview)

	return client.Do(ctx, req, nil)
}

func InSyncForce(client *github.Client, ctx context.Context, owner, repo, source_branch string, destination_branches []string, branchProtectionFlag string) {
	// In-Sync-Force (Rename-name-old-hash to create a new branch from source branch)
	// Check if Destination & Source Branch existes
	branches := append(destination_branches, source_branch)
	response, BranchNotFound := CheckBranchEval(ctx, client, owner, repo, branches)
	// if branchProtectionFlag == "disabled" {
	// 	for _, branch := range destination_branches {
	// 		RemoveBranchProtection(client, ctx, owner, repo, branch)
	// 	}
	// }
	fmt.Println("Response", response)
	if response {
		// Delete the existing branch
		for _, destination_branch := range destination_branches {
			DeleteBranch(client, ctx, owner, repo, destination_branch)
			fmt.Println("Destination Branch:", destination_branch)
			CreateBranch(client, ctx, owner, repo, source_branch, destination_branch)
		}
		return
	} else {
		panic("Branches not found:" + BranchNotFound)
	}

}

func Remove(client *github.Client, ctx context.Context, owner, repo string, veave *veave.Veaver, RepGetOptions github.RepositoryContentGetOptions) {
	// Direct get the file path and perform delete operation on that path
	// adding to .gitignore files list
	for index, _ := range veave.Rules {

	}
	RepGetOptions.Ref = branch
	FileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filepath, &RepGetOptions)
	if err != nil {
		fmt.Println("Error:", err)
	}
	var CommitMessage, SHA, Branch *string
	var Author *github.CommitAuthor
	var CommitAuthor *github.CommitAuthor

	DeleteFileOpts := &github.RepositoryContentFileOptions{
		Message:   CommitMessage,
		Content:   []byte{},
		SHA:       SHA,
		Branch:    Branch,
		Author:    Author,
		Committer: CommitAuthor,
	}
	DeleteFileOpts.SHA = FileContent.SHA
	fmt.Println(DeleteFileOpts)
	// FileContent.SHA
	DeleteContent, _, _, error := client.Repositories.DeleteFile(ctx, owner, repo, filepath, DeleteFileOpts)
	if error != nil {
		fmt.Println("Error:", error)
	}

}

func CheckIsContributor(client *github.Client, ctx context.Context, owner, repo string, userlist []string) (bool, error) {
	// Verify the github Users are Contributors
	var err error
	var isContributor bool
	for _, user := range userlist {
		isContributor, _, _ = client.Repositories.IsCollaborator(ctx, owner, repo, user)
		if isContributor == false {
			fmt.Println("2", isContributor)
			break
		}
		if err != nil {
			utils.HandleErr(err.Error())
		}
	}
	return isContributor, err
}

func GetRepoDetails(RepositoryFullName string) (string, string) {
	// Extract Owner + Repository details
	RepositoryFullNameData := strings.Split(RepositoryFullName, "/")
	org_name := RepositoryFullNameData[0]
	repo_name := RepositoryFullNameData[1]
	fmt.Println("OrgName: ", org_name, "RepoName: ", repo_name)
	return org_name, repo_name
}

func OpenPREval(context context.Context, client *github.Client, OrgName, RepoName string) (OpenPRFlag bool) {
	// // Creating struct to get OpenPRList
	OpenPRStruct := github.PullRequestListOptions{
		State: "open",
	}

	// Getting a list of OpenPRList
	OpenPRList, _, err := client.PullRequests.List(context, string(OrgName), string(RepoName), &OpenPRStruct)
	// fmt.Println("Open PR List: ", OpenPRList)
	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	// Check if there is any PR that is already opened for the same source destination branches to avoid recreation
	for index := range OpenPRList {
		fmt.Println(*OpenPRList[index].ID, *OpenPRList[index].Number, *OpenPRList[index].State, *OpenPRList[index].Title, utils.CheckVeavied(*OpenPRList[index].Title))
		if utils.CheckVeavied(*OpenPRList[index].Title) == true {
			isVeaveFlag := utils.CheckVeavied(*OpenPRList[index].Title)
			if isVeaveFlag {
				OpenPRFlag = false
			} else {
				OpenPRFlag = true
			}
			break
		}
	}
	return
}

func CheckBranchEval(context context.Context, client *github.Client, owner, repo string, branches []string) (BranchExistFlag bool, BranchNotFound string) {
	for _, branch := range branches {
		_, resp, err := client.Repositories.GetBranch(context, owner, repo, branch)
		if (err != nil) && (resp.StatusCode == 404) {
			BranchExistFlag = false
			BranchNotFound = branch
			fmt.Println("Branch Name:" + BranchNotFound + " not found.")
			fmt.Println("Error:", err)
			return
		} else {
			BranchExistFlag = true
			return
		}
	}
	return
}

func GetUser(ctx context.Context, gc *github.Client, user string) *github.User {
	// Get List of Github Users
	user_data, _, err := gc.Users.Get(ctx, user)
	if err != nil {
		log.Println(err)
	}
	return user_data
}

func GetReviewersDetails(reviewerpayload veave.Veaver) ([]string, []string) {
	// Get a list of reviewers + slack_id
	var ReviewerList, SlackIDList []string
	// for _, reviewer := range reviewerpayload.Reviewers {
	// 	ReviewerList = append(ReviewerList, reviewer.ID)
	// 	SlackIDList = append(SlackIDList, reviewer.Slackid)
	// }
	return ReviewerList, SlackIDList
}

func CreatePullRequest(ctx context.Context, gc *github.Client, owner, repo, title, HeadValue, BaseValue, BodyValue string, maintainer_can_modify bool, slackFlag bool, reviewers veave.Veaver, team_reviewers []string) {
	// Create Pull Request
	// BodyValue = GV_PRFLag + "BackMerge from " + HeadValue + " to " + BaseValue
	// title = "GITVEAVER " + "BackMerge from " + HeadValue + " to " + BaseValue

	pull := github.NewPullRequest{
		Title: &title,
		Body:  &BodyValue,
		Base:  &BaseValue,
		Head:  &HeadValue,
	}
	// fmt.Println("stuff:", owner, repo, title, HeadValue, BaseValue, BodyValue, maintainer_can_modify)
	prObject, _, err := gc.PullRequests.Create(ctx, owner, repo, &pull)
	if err != nil {
		log.Print("Error", err)
	}
	prNumber := *prObject.Number
	// fmt.Println("pr", prNumber)
	ReviewerIDList, SlackIDList := GetReviewersDetails(reviewers)
	fmt.Println(ReviewerIDList, SlackIDList)

	review_list := github.ReviewersRequest{
		Reviewers:     ReviewerIDList,
		TeamReviewers: team_reviewers,
	}
	fmt.Println(review_list)
	prObject, _, err = gc.PullRequests.RequestReviewers(ctx, owner, repo, prNumber, review_list)
	PRList := []string{}
	PRList = append(PRList, *prObject.HTMLURL)
	// fmt.Println("Request Reviewer", prObject.RequestedReviewers)
	// fmt.Println("Final Slack Payload:", SlackEnabledFlag)
	// SlackEnabledFlag = slackFlag
	// if err != nil {
	// 	log.Print("Error", err)
	// } else if SlackEnabledFlag == true {
	// 	fmt.Println(prObject.ID)
	// 	SlackNotifier("gitveaver", "GitVeaver", ":deepintent-2:", PRList, SlackIDList)
	// } else {
	// 	fmt.Println("NoSlackNotification")
	// }
}

func BackMerge(VeaverData veave.Veaver, VeaverRawPayload []byte) {
	// VeaverData.getConf(VeaverRawPayload)
	// VeaverPayload := GetVeaverData(VeaverRawPayload)
	// VeaverBranch := (VeaverPayload.Source)
	// ExecEvaluator := CheckExecutor(VeaverBranch, PushedBranch)
	// ReviewerList, _ := GetReviewersDetails(*VeaverPayload)
	// ContributorEval, _ := CheckIsContributor(client, context, Owner, Repo, ReviewerList)
	// PrefixCode := VeaverPayload.ConditionalPr.ConditionalPrTagPrefix
	// ConditionalPREval := ConditionalPRCheck(VeaverPayload.ConditionalPr.Enabled, ResponsePayload, PrefixCode)
	// CheckBranchEval := CheckBranchEval(context, client, Owner, Repo, VeaverPayload.Destination)
	// OpenPREval := OpenPREval(context, client, Owner, Repo)
	// fmt.Println("Evaulator: ", "ExecEval:", ExecEvaluator, "ContributorEval:", ContributorEval, "ConditinalPR:", ConditionalPREval, "CheckBranchEval:", CheckBranchEval, "OpenEval:", OpenPREval)

	// if (ExecEvaluator) && (ContributorEval) && (ConditionalPREval) && (CheckBranchEval) && (OpenPREval) {
	// 	PRCreation(context, client, *VeaverPayload, ResponsePayload, Owner, Repo)
	// } else {
	// 	fmt.Println("Unable to process the request.")
	// 	if !ExecEvaluator {
	// 		fmt.Println("Veaver Branch:", VeaverBranch, " not the same as ", PushedBranch)
	// 	} else if !ContributorEval {
	// 		fmt.Println("Repository doesnt contains the following as Contributors", ReviewerList)
	// 	} else if !ConditionalPREval {
	// 		fmt.Println("Looks like Commit Prefix didnt comply with the Conditiional PR Present in the configuration", ConditionalPREval)
	// 	} else if !CheckBranchEval {
	// 		fmt.Println("Branch specified in the configuration doesnt exist", VeaverPayload.Destination)
	// 	} else if !OpenPREval {
	// 		fmt.Println("PR Already created by GitVeaver Please close", OpenPREval)
	// 	}
	// }
}

func ConditionalPRCheck(ConditionPRFlag bool, resp veave.GithubPayload, PrefixCode string) (StatusCode bool) {
	// Check Condition PR Flag: if Condition PR Flag enabled will check if the commit has specific tag as a
	StatusCode = ConditionPRFlag
	if ConditionPRFlag {
		if strings.HasPrefix(resp.HeadCommit.Message, PrefixCode) {
			StatusCode = true
		} else {
			StatusCode = false
		}
	} else {
		StatusCode = true
	}
	return StatusCode
}

// func (c *veaver) getConf(yamlData []byte) *veaver {
// 	err := yaml.Unmarshal(yamlData, c)
// 	if err != nil {
// 		log.Fatalf("Unmarshal: %v", err)
// 	}
// 	return c
// }

// func PRCreation(context context.Context, client *github.Client, v veave.Veaver, r veave.GithubPayload, org_name, repo_name string) {
// 	ReviewersList := v
// 	TeamReviewersList := []string{}
// 	maintainer_can_modify := true
// 	fmt.Println("Data: -> ", org_name, repo_name, v.Title, v.Destination, v.Source, v.Description, maintainer_can_modify, ReviewersList, TeamReviewersList)

// 	fmt.Println("Final Evaulation IsVeaved:", EvalExpr)
// 	if EvalExpr == false {
// 		for index, _ := range v.Destination {
// 			fmt.Println("PR Creation Cycle: ", index)
// 			CreatePullRequest(context, client, org_name, repo_name, v.Title, v.Destination[index], v.Source, v.Description, maintainer_can_modify, v.SlackNotification.Enabled, ReviewersList, TeamReviewersList)
// 		}
// 	} else {
// 		return
// 	}
// 	for index, _ := range v.Destination {
// 		fmt.Println("PR Creation Cycle: ", index)
// 		CreatePullRequest(context, client, org_name, repo_name, v.Title, v.Destination[index], v.Source, v.Description, maintainer_can_modify, v.SlackNotification.Enabled, ReviewersList, TeamReviewersList)
// 	}
// }

func GetBranch(ref string) string {
	// Get Branch Details
	if strings.Contains(ref, "refs/heads/") {
		return string(ref[11:])
	}
	return ""
}
