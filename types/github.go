package types

// types used by github.com api

type GithubProjects []struct{}

type GithubProject struct{}

type GithubIssues []struct{}

type GithubIssue struct{}

type GithubRepositories []struct{}

type GithubRepositorie struct {
	ID            int    `json:"id,omitempty"`
	HtmlUrl       string `json:"html_url,omitempty"`
	DefaultBranch string `json:"default_branch,omitempty"`
	FullName      string `json:"full_name,omitempty"`
}

type GithubFile struct {
	Content     string `json:"content,omitempty"`
	DownloadUrl string `json:"download_url,omitempty"`
}
