package main

import (
	"github.com/hashicorp/go-plugin"
	"gitlab.com/jstang/rasic/types/plugins"
)

const baseURL = "https://gitlab.com"
const apiPath = "/api/v4/"

const OK = "200 OK"

type SourceGithub struct{}


func main() {
	github := &SourceGithub{}

	var pluginMap = map[string]plugin.Plugin{
		"source_github": &plugins.SourcePlugin{Impl: github},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
