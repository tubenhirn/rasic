package core

import (
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

func DispensePlugins(pluginList []types.RasicPlugin, logger hclog.Logger) (plugins.API, plugins.Reporter, []*plugin.Client) {
	var returnAPIPlugin plugins.API
	var returnReporterPlugin plugins.Reporter

	// collect all clients to kill them after use
	// types does not matter here
	var clientList []*plugin.Client

	for _, currentPlugin := range pluginList {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: currentPlugin.PluginConfig,
			Plugins:         currentPlugin.PluginMap,
			Cmd:             exec.Command(currentPlugin.PluginHome + currentPlugin.PluginPath + "/" + currentPlugin.PluginName),
			Logger:          logger,
		})

		rpcClient, err := client.Client()
		if err != nil {
			pterm.Error.Println(err)
		}

		raw, dispenseErr := rpcClient.Dispense(currentPlugin.PluginName)
		if dispenseErr != nil {
			pterm.Error.Println(dispenseErr)
		}
		switch currentPlugin.PluginPath {
		case "api":
			plug := raw.(plugins.API)
			returnAPIPlugin = plug
		case "reporter":
			plug := raw.(plugins.Reporter)
			returnReporterPlugin = plug
		default:
			pterm.Warning.Println("plugin could not be loaded")
		}
		clientList = append(clientList, client)
	}

	return returnAPIPlugin, returnReporterPlugin, clientList
}
