package types

import (
	"net/http"

	"github.com/hashicorp/go-plugin"
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

type RasicPlugin struct {
	PluginPath   string
	PluginName   string
	PluginConfig plugin.HandshakeConfig
	PluginMap    map[string]plugin.Plugin
}
