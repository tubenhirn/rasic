package core

import (
	"os/exec"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

func DispensePlugins(pluginList []types.RasicPlugin, logger hclog.Logger) (plugins.Source, plugins.Reporter, []*plugin.Client) {
	var returnAPIPlugin plugins.Source
	var returnReporterPlugin plugins.Reporter

	// collect all clients to kill them after use
	// types does not matter here
	var clientList []*plugin.Client

	for _, currentPlugin := range pluginList {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: currentPlugin.PluginConfig,
			Plugins:         currentPlugin.PluginMap,
			// TODO: use again when plugin is reworked
			// Cmd:             exec.Command(currentPlugin.PluginHome + currentPlugin.PluginPath + "/" + currentPlugin.PluginName),
			Cmd:             exec.Command(currentPlugin.PluginHome + "/bin/" + currentPlugin.PluginName),
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
		// TODO: remove this when plugin is reworked
		switch strings.Split(currentPlugin.PluginName, "_")[0] {
		case "source":
			plug := raw.(plugins.Source)
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
