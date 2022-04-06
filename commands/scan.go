package commands

import (
	"net/http"
	"os"
	"os/exec"
	"strconv"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"gitlab.com/jstang/rasic/scan"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

var (
	projectFlag = cli.StringFlag{
		Name:        "project",
		Aliases:     []string{"group"},
		Usage:       "a project or group id to scan",
		EnvVars:     []string{},
		FilePath:    "",
		Required:    true,
		Hidden:      false,
		TakesFile:   false,
		Value:       "",
		DefaultText: "",
		Destination: new(string),
		HasBeenSet:  false,
	}
	tokenFlag = cli.StringFlag{
		Name:        "token",
		Aliases:     []string{},
		Usage:       "a oauth token",
		EnvVars:     []string{"GITLAB_TOKEN"},
		FilePath:    "",
		Required:    true,
		Hidden:      false,
		TakesFile:   false,
		Value:       "",
		DefaultText: "",
		Destination: new(string),
		HasBeenSet:  false,
	}
	ignoreFileFlag = cli.StringFlag{
		Name:        "ignorefile",
		Aliases:     []string{},
		Usage:       "specify a ignorefile. set ignorefile flag or use CVE_IGNORE_FILE env. (default .trivyignore)",
		EnvVars:     []string{"CVE_IGNORE_FILE"},
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		TakesFile:   false,
		Value:       ".trivyignore",
		DefaultText: "",
		Destination: new(string),
		HasBeenSet:  false,
	}
)

// scan a project for cve's
func Scan() *cli.Command {
	return &cli.Command{
		Name:        "scan",
		Aliases:     []string{"s"},
		Usage:       "scan project for cve's or config flaws",
		UsageText:   "",
		Description: "",
		ArgsUsage:   "",
		Category:    "",
		BashComplete: func(*cli.Context) {
		},
		Before: func(*cli.Context) error {
			return nil
		},
		After: func(*cli.Context) error {
			return nil
		},
		Action: func(c *cli.Context) error {
			sourceName := c.String("source")
			reporterName := c.String("reporter")
			projectId := c.String("project")
			authToken := c.String("token")
			ignoreFileName := c.String("ignorefile")

			var handshakeConfig = plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "API_PLUGIN",
				MagicCookieValue: "allow",
			}

			var pluginMap = map[string]plugin.Plugin{
				"gitlab": &plugins.ApiPlugin{},
				"github": &plugins.ApiPlugin{},
				"trivy":  &plugins.ApiPlugin{},
			}

			httpClient := &http.Client{}
			logger := hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stdout,
				Level:  hclog.Error,
			})

			// load all plugins required for this command
			plugins, clients := dispensePlugins([]string{sourceName, reporterName}, handshakeConfig, pluginMap, logger)

			for _, pluginClient := range clients {
				defer pluginClient.Kill()
			}

			pterm.Info.Println("scan for cve's")

			projects := plugins[sourceName].GetProjects(httpClient, projectId, authToken)
			if len(projects) < 1 {
				pterm.Info.Println("no projects found in group " + projectId + "(maybe it is a project?)")

				singleProject := plugins[sourceName].GetProject(httpClient, projectId, authToken)
				var currentProject types.RasicProject
				currentProject.Id = singleProject.Id
				currentProject.WebUrl = singleProject.WebUrl
				currentProject.DefaultBranch = singleProject.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				var issues []types.RasicIssue
				issues = plugins[sourceName].GetIssues(httpClient, strconv.Itoa(singleProject.Id), authToken)
				pterm.Info.Printfln("scan: " + currentProject.WebUrl)

				issues, err := scan.Scanner(httpClient, plugins[sourceName], plugins[reporterName], currentProject, authToken, issues)
				if err != nil {
					pterm.Error.Println(err)
				}

				for _, issue := range issues {
					plugins[reporterName].CreateIssue(httpClient, strconv.Itoa(singleProject.Id), authToken, issue)
					pterm.Info.Println("new issue opened for " + issue.Title)
				}

				return nil
			}
			pterm.Info.Println(strconv.Itoa(len(projects)) + " projects found in group " + projectId)
			for _, project := range projects {
				var issues []types.RasicIssue

				var currentProject types.RasicProject
				currentProject.Id = project.Id
				currentProject.WebUrl = project.WebUrl
				currentProject.DefaultBranch = project.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				issues = plugins[sourceName].GetIssues(httpClient, strconv.Itoa(project.Id), authToken)
				pterm.Info.Println("scan: " + project.WebUrl)

				issues, err := scan.Scanner(httpClient, plugins[sourceName], plugins[reporterName], currentProject, authToken, issues)
				if err != nil {
					pterm.Error.Println(err)
				}

				// TODO issue can be created elsewhere
				// if we push them to jira we need a different target (project.Id)
				for _, issue := range issues {
					plugins[reporterName].CreateIssue(httpClient, strconv.Itoa(project.Id), authToken, issue)
					pterm.Info.Println("new issue opened for " + issue.Title)
				}
			}

			return nil
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Subcommands:            []*cli.Command{},
		Flags:                  []cli.Flag{&projectFlag, &tokenFlag, &ignoreFileFlag},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

}

// dispense a map of plugins and list of client selected by name
func dispensePlugins(pluginNameList []string, config plugin.HandshakeConfig, pluginMap map[string]plugin.Plugin, logger hclog.Logger) (map[string]plugins.Api, []*plugin.Client) {

	pluginList := make(map[string]plugins.Api)
	var clientList []*plugin.Client

	for _, pluginName := range pluginNameList {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: config,
			Plugins:         pluginMap,
			Cmd:             exec.Command("./plugins/api/" + pluginName),
			Logger:          logger,
		})

		rpcClient, err := client.Client()
		if err != nil {
			pterm.Error.Println(err)
		}

		raw, reporterErr := rpcClient.Dispense(pluginName)
		if reporterErr != nil {
			pterm.Error.Println(reporterErr)
		}
		plug := raw.(plugins.Api)
		pluginList[pluginName] = plug
		clientList = append(clientList, client)

	}
	return pluginList, clientList
}
