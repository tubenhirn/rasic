package types

import (
	"net/http"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type RasicIssue struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string
}

type RasicProject struct {
	Id             int
	WebUrl         string
	DefaultBranch  string
	IgnoreFileName string
}
