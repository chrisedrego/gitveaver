package veave

import (
	"time"
)

type Veaver struct {
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
		BranchProtection    string   `yaml:"branch_protection,omitempty"`
		DestinationBranches []string `yaml:"destination_branches,omitempty"`
		Message             string   `yaml:"message,omitempty"`
		AuthorID            string   `yaml:"author_id,omitempty"`
		AuthorEmail         string   `yaml:"author_email,omitempty"`
		Path                []string `yaml:"path,omitempty"`
	} `yaml:"rules"`
}

type GithubPayload struct {
	Ref        string `yaml:"ref"`
	Before     string `yaml:"before"`
	After      string `yaml:"after"`
	Repository struct {
		ID        int    `yaml:"id"`
		NodeID    string `yaml:"node_id"`
		Name      string `yaml:"name"`
		Full_Name string `yaml:"full_name"`
		Private   bool   `yaml:"private"`
		Owner     struct {
			Name              string `yaml:"name"`
			Email             string `yaml:"email"`
			Login             string `yaml:"login"`
			ID                int    `yaml:"id"`
			NodeID            string `yaml:"node_id"`
			AvatarURL         string `yaml:"avatar_url"`
			GravatarID        string `yaml:"gravatar_id"`
			URL               string `yaml:"url"`
			HTMLURL           string `yaml:"html_url"`
			FollowersURL      string `yaml:"followers_url"`
			FollowingURL      string `yaml:"following_url"`
			GistsURL          string `yaml:"gists_url"`
			StarredURL        string `yaml:"starred_url"`
			SubscriptionsURL  string `yaml:"subscriptions_url"`
			OrganizationsURL  string `yaml:"organizations_url"`
			ReposURL          string `yaml:"repos_url"`
			EventsURL         string `yaml:"events_url"`
			ReceivedEventsURL string `yaml:"received_events_url"`
			Type              string `yaml:"type"`
			SiteAdmin         bool   `yaml:"site_admin"`
		} `yaml:"owner"`
		HTMLURL          string      `yaml:"html_url"`
		Description      string      `yaml:"description"`
		Fork             bool        `yaml:"fork"`
		URL              string      `yaml:"url"`
		ForksURL         string      `yaml:"forks_url"`
		KeysURL          string      `yaml:"keys_url"`
		CollaboratorsURL string      `yaml:"collaborators_url"`
		TeamsURL         string      `yaml:"teams_url"`
		HooksURL         string      `yaml:"hooks_url"`
		IssueEventsURL   string      `yaml:"issue_events_url"`
		EventsURL        string      `yaml:"events_url"`
		AssigneesURL     string      `yaml:"assignees_url"`
		BranchesURL      string      `yaml:"branches_url"`
		TagsURL          string      `yaml:"tags_url"`
		BlobsURL         string      `yaml:"blobs_url"`
		GitTagsURL       string      `yaml:"git_tags_url"`
		GitRefsURL       string      `yaml:"git_refs_url"`
		TreesURL         string      `yaml:"trees_url"`
		StatusesURL      string      `yaml:"statuses_url"`
		LanguagesURL     string      `yaml:"languages_url"`
		StargazersURL    string      `yaml:"stargazers_url"`
		ContributorsURL  string      `yaml:"contributors_url"`
		SubscribersURL   string      `yaml:"subscribers_url"`
		SubscriptionURL  string      `yaml:"subscription_url"`
		CommitsURL       string      `yaml:"commits_url"`
		GitCommitsURL    string      `yaml:"git_commits_url"`
		CommentsURL      string      `yaml:"comments_url"`
		IssueCommentURL  string      `yaml:"issue_comment_url"`
		ContentsURL      string      `yaml:"contents_url"`
		CompareURL       string      `yaml:"compare_url"`
		MergesURL        string      `yaml:"merges_url"`
		ArchiveURL       string      `yaml:"archive_url"`
		DownloadsURL     string      `yaml:"downloads_url"`
		IssuesURL        string      `yaml:"issues_url"`
		PullsURL         string      `yaml:"pulls_url"`
		MilestonesURL    string      `yaml:"milestones_url"`
		NotificationsURL string      `yaml:"notifications_url"`
		LabelsURL        string      `yaml:"labels_url"`
		ReleasesURL      string      `yaml:"releases_url"`
		DeploymentsURL   string      `yaml:"deployments_url"`
		CreatedAt        int         `yaml:"created_at"`
		UpdatedAt        time.Time   `yaml:"updated_at"`
		PushedAt         int         `yaml:"pushed_at"`
		GitURL           string      `yaml:"git_url"`
		SSHURL           string      `yaml:"ssh_url"`
		CloneURL         string      `yaml:"clone_url"`
		SvnURL           string      `yaml:"svn_url"`
		Homepage         interface{} `yaml:"homepage"`
		Size             int         `yaml:"size"`
		StargazersCount  int         `yaml:"stargazers_count"`
		WatchersCount    int         `yaml:"watchers_count"`
		Language         string      `yaml:"language"`
		HasIssues        bool        `yaml:"has_issues"`
		HasProjects      bool        `yaml:"has_projects"`
		HasDownloads     bool        `yaml:"has_downloads"`
		HasWiki          bool        `yaml:"has_wiki"`
		HasPages         bool        `yaml:"has_pages"`
		ForksCount       int         `yaml:"forks_count"`
		MirrorURL        interface{} `yaml:"mirror_url"`
		Archived         bool        `yaml:"archived"`
		Disabled         bool        `yaml:"disabled"`
		OpenIssuesCount  int         `yaml:"open_issues_count"`
		License          struct {
			Key    string `yaml:"key"`
			Name   string `yaml:"name"`
			SpdxID string `yaml:"spdx_id"`
			URL    string `yaml:"url"`
			NodeID string `yaml:"node_id"`
		} `yaml:"license"`
		AllowForking             bool          `yaml:"allow_forking"`
		IsTemplate               bool          `yaml:"is_template"`
		WebCommitSignoffRequired bool          `yaml:"web_commit_signoff_required"`
		Topics                   []interface{} `yaml:"topics"`
		Visibility               string        `yaml:"visibility"`
		Forks                    int           `yaml:"forks"`
		OpenIssues               int           `yaml:"open_issues"`
		Watchers                 int           `yaml:"watchers"`
		DefaultBranch            string        `yaml:"default_branch"`
		Stargazers               int           `yaml:"stargazers"`
		MasterBranch             string        `yaml:"master_branch"`
	} `yaml:"repository"`
	Pusher struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
	} `yaml:"pusher"`
	Sender struct {
		Login             string `yaml:"login"`
		ID                int    `yaml:"id"`
		NodeID            string `yaml:"node_id"`
		AvatarURL         string `yaml:"avatar_url"`
		GravatarID        string `yaml:"gravatar_id"`
		URL               string `yaml:"url"`
		HTMLURL           string `yaml:"html_url"`
		FollowersURL      string `yaml:"followers_url"`
		FollowingURL      string `yaml:"following_url"`
		GistsURL          string `yaml:"gists_url"`
		StarredURL        string `yaml:"starred_url"`
		SubscriptionsURL  string `yaml:"subscriptions_url"`
		OrganizationsURL  string `yaml:"organizations_url"`
		ReposURL          string `yaml:"repos_url"`
		EventsURL         string `yaml:"events_url"`
		ReceivedEventsURL string `yaml:"received_events_url"`
		Type              string `yaml:"type"`
		SiteAdmin         bool   `yaml:"site_admin"`
	} `yaml:"sender"`
	Created bool        `yaml:"created"`
	Deleted bool        `yaml:"deleted"`
	Forced  bool        `yaml:"forced"`
	BaseRef interface{} `yaml:"base_ref"`
	Compare string      `yaml:"compare"`
	Commits []struct {
		ID        string    `yaml:"id"`
		TreeID    string    `yaml:"tree_id"`
		Distinct  bool      `yaml:"distinct"`
		Message   string    `yaml:"message"`
		Timestamp time.Time `yaml:"timestamp"`
		URL       string    `yaml:"url"`
		Author    struct {
			Name     string `yaml:"name"`
			Email    string `yaml:"email"`
			Username string `yaml:"username"`
		} `yaml:"author"`
		Committer struct {
			Name     string `yaml:"name"`
			Email    string `yaml:"email"`
			Username string `yaml:"username"`
		} `yaml:"committer"`
		Added    []interface{} `yaml:"added"`
		Removed  []interface{} `yaml:"removed"`
		Modified []interface{} `yaml:"modified"`
	} `yaml:"commits"`
	HeadCommit struct {
		ID        string    `yaml:"id"`
		TreeID    string    `yaml:"tree_id"`
		Distinct  bool      `yaml:"distinct"`
		Message   string    `yaml:"message"`
		Timestamp time.Time `yaml:"timestamp"`
		URL       string    `yaml:"url"`
		Author    struct {
			Name     string `yaml:"name"`
			Email    string `yaml:"email"`
			Username string `yaml:"username"`
		} `yaml:"author"`
		Committer struct {
			Name     string `yaml:"name"`
			Email    string `yaml:"email"`
			Username string `yaml:"username"`
		} `yaml:"committer"`
		Added    []interface{} `yaml:"added"`
		Removed  []interface{} `yaml:"removed"`
		Modified []interface{} `yaml:"modified"`
	} `yaml:"head_commit"`
}
