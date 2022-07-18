package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chrisedrego/gitveaver/utils"
	"github.com/google/go-github/github"
)

type GithubPayload struct {
	// Github Payload
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Name              string `json:"name"`
			Email             string `json:"email"`
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HTMLURL          string      `json:"html_url"`
		Description      string      `json:"description"`
		Fork             bool        `json:"fork"`
		URL              string      `json:"url"`
		ForksURL         string      `json:"forks_url"`
		KeysURL          string      `json:"keys_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		TeamsURL         string      `json:"teams_url"`
		HooksURL         string      `json:"hooks_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		EventsURL        string      `json:"events_url"`
		AssigneesURL     string      `json:"assignees_url"`
		BranchesURL      string      `json:"branches_url"`
		TagsURL          string      `json:"tags_url"`
		BlobsURL         string      `json:"blobs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		TreesURL         string      `json:"trees_url"`
		StatusesURL      string      `json:"statuses_url"`
		LanguagesURL     string      `json:"languages_url"`
		StargazersURL    string      `json:"stargazers_url"`
		ContributorsURL  string      `json:"contributors_url"`
		SubscribersURL   string      `json:"subscribers_url"`
		SubscriptionURL  string      `json:"subscription_url"`
		CommitsURL       string      `json:"commits_url"`
		GitCommitsURL    string      `json:"git_commits_url"`
		CommentsURL      string      `json:"comments_url"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		ContentsURL      string      `json:"contents_url"`
		CompareURL       string      `json:"compare_url"`
		MergesURL        string      `json:"merges_url"`
		ArchiveURL       string      `json:"archive_url"`
		DownloadsURL     string      `json:"downloads_url"`
		IssuesURL        string      `json:"issues_url"`
		PullsURL         string      `json:"pulls_url"`
		MilestonesURL    string      `json:"milestones_url"`
		NotificationsURL string      `json:"notifications_url"`
		LabelsURL        string      `json:"labels_url"`
		ReleasesURL      string      `json:"releases_url"`
		DeploymentsURL   string      `json:"deployments_url"`
		CreatedAt        int         `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		PushedAt         int         `json:"pushed_at"`
		GitURL           string      `json:"git_url"`
		SSHURL           string      `json:"ssh_url"`
		CloneURL         string      `json:"clone_url"`
		SvnURL           string      `json:"svn_url"`
		Homepage         interface{} `json:"homepage"`
		Size             int         `json:"size"`
		StargazersCount  int         `json:"stargazers_count"`
		WatchersCount    int         `json:"watchers_count"`
		Language         interface{} `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasProjects      bool        `json:"has_projects"`
		HasDownloads     bool        `json:"has_downloads"`
		HasWiki          bool        `json:"has_wiki"`
		HasPages         bool        `json:"has_pages"`
		ForksCount       int         `json:"forks_count"`
		MirrorURL        interface{} `json:"mirror_url"`
		Archived         bool        `json:"archived"`
		Disabled         bool        `json:"disabled"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		License          struct {
			Key    string `json:"key"`
			Name   string `json:"name"`
			SpdxID string `json:"spdx_id"`
			URL    string `json:"url"`
			NodeID string `json:"node_id"`
		} `json:"license"`
		AllowForking  bool          `json:"allow_forking"`
		IsTemplate    bool          `json:"is_template"`
		Topics        []interface{} `json:"topics"`
		Visibility    string        `json:"visibility"`
		Forks         int           `json:"forks"`
		OpenIssues    int           `json:"open_issues"`
		Watchers      int           `json:"watchers"`
		DefaultBranch string        `json:"default_branch"`
		Stargazers    int           `json:"stargazers"`
		MasterBranch  string        `json:"master_branch"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Sender struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
	Created bool        `json:"created"`
	Deleted bool        `json:"deleted"`
	Forced  bool        `json:"forced"`
	BaseRef interface{} `json:"base_ref"`
	Compare string      `json:"compare"`
	Commits []struct {
		ID        string    `json:"id"`
		TreeID    string    `json:"tree_id"`
		Distinct  bool      `json:"distinct"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []interface{} `json:"added"`
		Removed  []interface{} `json:"removed"`
		Modified []string      `json:"modified"`
	} `json:"commits"`
	HeadCommit struct {
		ID        string    `json:"id"`
		TreeID    string    `json:"tree_id"`
		Distinct  bool      `json:"distinct"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []interface{} `json:"added"`
		Removed  []interface{} `json:"removed"`
		Modified []string      `json:"modified"`
	} `json:"head_commit"`
}

type Veaver struct {
	// Veaver Payload
	Rules []struct {
		Name              string   `yaml:"name"`
		Mode              string   `yaml:"mode"`
		SourceBranch      string   `yaml:"source_branch"`
		DestinationBranch []string `yaml:"destination_branch,omitempty"`
		Conditional       struct {
			Enabled   bool   `yaml:"enabled"`
			TagPrefix string `yaml:"tag_prefix"`
		} `yaml:"conditional,omitempty"`
		Reviewers         []string `yaml:"reviewers,omitempty"`
		Title             string   `yaml:"title,omitempty"`
		Description       string   `yaml:"description,omitempty"`
		SlackNotification struct {
			Enabled bool     `yaml:"enabled"`
			SlackID []string `yaml:"slack_id"`
		} `yaml:"slack_notification,omitempty"`
		DestinationRepo  string `yaml:"destination_repo,omitempty"`
		DestinationRules []struct {
			Name        string   `yaml:"name"`
			ExcludeDir  []string `yaml:"exclude dir"`
			ExcludeFile []string `yaml:"exclude_file"`
		} `yaml:"destination_rules,omitempty"`
	} `yaml:"rules"`
}

type error interface {
	Error() string
}

// Global FLAG
// var SlackEnabledFlag bool
// var GV_PRFLag string = "GITVEAVER"
var configFile string
var githubToken string

func handleErr(err string) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
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
			handleErr(err.Error())
		}
	}
	return isContributor, err
}

func getRawVeaver(client *github.Client, ctx context.Context, owner, repo, filepath, branch string, RepGetOptions github.RepositoryContentGetOptions) []byte {
	// Get Contents of GitVeaver Configuration
	FileContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filepath, &RepGetOptions)
	if err != nil {
		handleErr(err.Error())
	}
	rawDecodedData, err := base64.StdEncoding.DecodeString(*FileContent.Content)
	if err != nil {
		panic(err)
	}
	return rawDecodedData
}

// func GetVeaverData(rawdata []byte) *veaver {
// 	// Get Veaver Data Struct
// 	// var data *veaver
// 	// data_err := yaml.Unmarshal(rawdata, &data)

// 	// if data_err != nil {
// 	// 	log.Fatal(data_err)
// 	// }
// 	// return data
// }

func GetRepoDetails(RepositoryFullName string) (string, string) {
	// Extract Owner + Repository details
	RepositoryFullNameData := strings.Split(RepositoryFullName, "/")
	org_name := RepositoryFullNameData[0]
	repo_name := RepositoryFullNameData[1]
	fmt.Println("OrgName: ", org_name, "RepoName: ", repo_name)
	return org_name, repo_name
}

func constructSlackMessage(prList, listSlackUser []string) (slackMessage string) {
	// Construct Slack Message
	var SlackUserId string
	SlackUserId = strings.Join(listSlackUser, " ")
	fmt.Println(SlackUserId)
	SlackMessageText := "\nWe need your attention, could you quickly reviews these Pull Request assigned to you."
	var prTextPayload string
	for _, val := range prList {
		prTextPayload = val + "\n"
		fmt.Println(prTextPayload)
	}
	slackMessage = "Hey! :wave:" + SlackUserId + SlackMessageText + "\n" + prTextPayload
	fmt.Println("Slack Message:", slackMessage)
	return slackMessage
}

func SlackNotifier(channel, username, icon_emoji string, prList, slackUserList []string) {
	// Send Slack Notification
	slackUrl := os.Getenv("SLACK_URI")
	fmt.Println("Slack-URI Details", slackUrl)
	textPayload := constructSlackMessage(prList, slackUserList)
	payload := map[string]interface{}{"channel": channel, "username": username, "icon_emoji": icon_emoji, "text": textPayload}
	// fmt.Println(payload)
	postBody, _ := json.Marshal(payload)
	// fmt.Println(postBody)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(slackUrl, "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(json.NewDecoder(resp.Body))
	fmt.Println(resp.StatusCode)
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
		fmt.Println(*OpenPRList[index].ID, *OpenPRList[index].Number, *OpenPRList[index].State, *OpenPRList[index].Title, isVeavied(*OpenPRList[index].Title))
		if isVeavied(*OpenPRList[index].Title) == true {
			isVeaveFlag := isVeavied(*OpenPRList[index].Title)
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

func CheckBranchEval(context context.Context, client *github.Client, owner, repo string, branches []string) (BranchExistFlag bool) {
	for _, branch := range branches {
		_, resp, err := client.Repositories.GetBranch(context, owner, repo, branch)
		if (err != nil) && (resp.StatusCode == 404) {
			BranchExistFlag = false
			return
		} else {
			BranchExistFlag = true
			return
		}
	}
	return
}

// func PRCreation(context context.Context, client *github.Client, v veaver, r GithubPayload, org_name, repo_name string) {
// // Creating Github Context
// context := context.Background()
// githubtoken := os.Getenv("GITHUBTOKEN")
// fmt.Println("Github Auth Details", githubtoken)

// tokenService := oauth2.StaticTokenSource(
// 	&oauth2.Token{AccessToken: githubtoken},
// )
// tokenClient := oauth2.NewClient(context, tokenService)

// // Intiating Github Client
// client := github.NewClient(tokenClient)

// ReviewersList := v
// TeamReviewersList := []string{}
// maintainer_can_modify := true
// fmt.Println("Data: -> ", org_name, repo_name, v.Title, v.Destination, v.Source, v.Description, maintainer_can_modify, ReviewersList, TeamReviewersList)

// fmt.Println("Final Evaulation IsVeaved:", EvalExpr)
// if EvalExpr == false {
// 	for index, _ := range v.Destination {
// 		fmt.Println("PR Creation Cycle: ", index)
// 		CreatePullRequest(context, client, org_name, repo_name, v.Title, v.Destination[index], v.Source, v.Description, maintainer_can_modify, v.SlackNotification.Enabled, ReviewersList, TeamReviewersList)
// 	}
// } else {
// 	return
// }
// for index, _ := range v.Destination {
// 	fmt.Println("PR Creation Cycle: ", index)
// 	CreatePullRequest(context, client, org_name, repo_name, v.Title, v.Destination[index], v.Source, v.Description, maintainer_can_modify, v.SlackNotification.Enabled, ReviewersList, TeamReviewersList)
// }
// }

func ConditionalPRCheck(ConditionPRFlag bool, resp GithubPayload, PrefixCode string) (StatusCode bool) {
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

func BackMerge(VeaverData Veaver, VeaverRawPayload []byte) {
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

func GetBranch(ref string) string {
	// Get Branch Details
	if strings.Contains(ref, "refs/heads/") {
		return string(ref[11:])
	}
	return ""
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

func (veave *Veaver) EvalChecker() {
	for index, _ := range veave.Rules {
		mode := veave.Rules[index].Mode
		switch mode {
		case "backmerge":
			fmt.Println("mode: backmerge")
		case "sync":
			fmt.Println("mode: sync")
		}
	}
}

func main() {
	utils.FlagCheck()
	http.HandleFunc("/", RequestHandler)
	log.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func isVeavied(data string) bool {
	var GV_PRFLag string
	return strings.HasPrefix(data, GV_PRFLag)
}

func GetUser(ctx context.Context, gc *github.Client, user string) *github.User {
	// Get List of Github Users
	user_data, _, err := gc.Users.Get(ctx, user)
	if err != nil {
		log.Println(err)
	}
	return user_data
}

func GetReviewersDetails(reviewerpayload Veaver) ([]string, []string) {
	// Get a list of reviewers + slack_id
	var ReviewerList, SlackIDList []string
	// for _, reviewer := range reviewerpayload.Reviewers {
	// 	ReviewerList = append(ReviewerList, reviewer.ID)
	// 	SlackIDList = append(SlackIDList, reviewer.Slackid)
	// }
	return ReviewerList, SlackIDList
}

func CreatePullRequest(ctx context.Context, gc *github.Client, owner, repo, title, HeadValue, BaseValue, BodyValue string, maintainer_can_modify bool, slackFlag bool, reviewers Veaver, team_reviewers []string) {
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
