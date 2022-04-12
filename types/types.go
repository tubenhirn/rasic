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

type RasicRepository struct {
	Id  int
	Tag RasicTag
}

type RasicTag struct {
	Location string
	Name     string
	Path     string
}

type RasicPlugin struct {
	PluginHome   string
	PluginPath   string
	PluginName   string
	PluginConfig plugin.HandshakeConfig
	PluginMap    map[string]plugin.Plugin
}
