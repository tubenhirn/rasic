package commands

import (
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

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
		EnvVars:     []string{"GITLAB_TOKEN", "RASIC_TOKEN"},
		FilePath:    "",
		Required:    true,
		Hidden:      false,
		TakesFile:   false,
		Value:       "",
		DefaultText: "",
		Destination: new(string),
		HasBeenSet:  false,
	}
	userNameFlag = cli.StringFlag{
		Name:        "user",
		Aliases:     []string{},
		Usage:       "a username used by trivy image scanning",
		EnvVars:     []string{"RASIC_USERNAME"},
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
		Usage:       "specify a cve ignorefile",
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
	containerScannerFlag = cli.BoolFlag{
		Name:        "container",
		Aliases:     []string{},
		Usage:       "enable container scanning",
		EnvVars:     []string{},
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		Value:       false,
		DefaultText: "",
		Destination: new(bool),
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
			pluginHome := c.String("pluginhome")
			projectId := c.String("project")
			userName := c.String("user")
			authToken := c.String("token")
			ignoreFileName := c.String("ignorefile")

			scanContainers := c.Bool("container")

			var apihandshakeConfig = plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "API_PLUGIN",
				MagicCookieValue: "allow",
			}
			var reporterhandshakeConfig = plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "REPORTER_PLUGIN",
				MagicCookieValue: "allow",
			}

			var apiPluginMap = map[string]plugin.Plugin{
				"gitlab": &plugins.ApiPlugin{},
			}
			var reporterPluginMap = map[string]plugin.Plugin{
				"gitlab": &plugins.ReporterPlugin{},
			}

			httpClient := &http.Client{}

			logger := hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stdout,
				Level:  hclog.Error,
			})

			pluginData := []types.RasicPlugin{
				{
					PluginPath:   "api",
					PluginHome:   pluginHome,
					PluginName:   sourceName,
					PluginConfig: apihandshakeConfig,
					PluginMap:    apiPluginMap,
				},
				{
					PluginPath:   "reporter",
					PluginHome:   pluginHome,
					PluginName:   reporterName,
					PluginConfig: reporterhandshakeConfig,
					PluginMap:    reporterPluginMap,
				},
			}

			// load all plugins required for this command
			apiPlugin, reporterPlugin, clients := dispensePlugins(pluginData, logger)

			for _, pluginClient := range clients {
				defer pluginClient.Kill()
			}

			pterm.Info.Println("scan for cve's")

			projects := apiPlugin.GetProjects(httpClient, projectId, authToken)
			if len(projects) < 1 {
				pterm.Info.Println("no projects found in group " + projectId + "(maybe it is a project?)")

				singleProject := apiPlugin.GetProject(httpClient, projectId, authToken)
				var currentProject types.RasicProject
				currentProject.Id = singleProject.Id
				currentProject.WebUrl = singleProject.WebUrl
				currentProject.DefaultBranch = singleProject.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				var newIssues []types.RasicIssue

				// scan current projects repositry (fs)
				pterm.Info.Printfln("scan repository: " + currentProject.WebUrl)
				tmpIssues, err := scan.RepositoryScanner(httpClient, apiPlugin, currentProject, authToken, newIssues)
				newIssues = append(newIssues, tmpIssues...)
				if err != nil {
					pterm.Error.Println(err)
				}

				// scan the project contaienr registry if enabled
				if scanContainers == true {
					newIssues = containerRegistryScan(httpClient, apiPlugin, currentProject, userName, authToken, newIssues)
				}

				openNewIssues(httpClient, reporterPlugin, currentProject, newIssues, authToken)

				return nil
			}

			// scan a group
			pterm.Info.Println(strconv.Itoa(len(projects)) + " projects found in group " + projectId)
			for _, project := range projects {

				var newIssues []types.RasicIssue

				var currentProject types.RasicProject
				currentProject.Id = project.Id
				currentProject.WebUrl = project.WebUrl
				currentProject.DefaultBranch = project.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				pterm.Info.Println("scan: " + project.WebUrl)

				tmpIssues, err := scan.RepositoryScanner(httpClient, apiPlugin, currentProject, authToken, newIssues)
				newIssues = append(newIssues, tmpIssues...)
				if err != nil {
					pterm.Error.Println(err)
				}

				// scan the project contaienr registry if enabled
				if scanContainers == true {
					newIssues = containerRegistryScan(httpClient, apiPlugin, currentProject, userName, authToken, newIssues)
				}

				openNewIssues(httpClient, reporterPlugin, currentProject, newIssues, authToken)
			}

			return nil
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Subcommands:            []*cli.Command{},
		Flags:                  []cli.Flag{&projectFlag, &tokenFlag, &userNameFlag, &ignoreFileFlag, &containerScannerFlag},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

}

// dispense a map of plugins (source, scanner, reporter) and list of client selected by name
// TODO maybe a bit diry - needs rework
func dispensePlugins(pluginList []types.RasicPlugin, logger hclog.Logger) (plugins.Api, plugins.Reporter, []*plugin.Client) {

	var returnApiPlugin plugins.Api
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
			plug := raw.(plugins.Api)
			returnApiPlugin = plug
		case "reporter":
			plug := raw.(plugins.Reporter)
			returnReporterPlugin = plug
		default:
			pterm.Warning.Println("plugin could not be loaded")
		}
		clientList = append(clientList, client)
	}

	return returnApiPlugin, returnReporterPlugin, clientList
}

// open new issues using the current reporter
func openNewIssues(httpClient types.HttpClient, reporterPlugin plugins.Reporter, project types.RasicProject, newIssues []types.RasicIssue, authToken string) {

	// get all issues for current project
	var projectIssues []types.RasicIssue
	projectIssues = reporterPlugin.GetIssues(httpClient, strconv.Itoa(project.Id), authToken)

	// check newIssues against projectIssues
	// if the issue does not exist in State="opened", create it with the current reporter
	for _, newIssue := range newIssues {
		issueExists := false
		for _, openIssue := range projectIssues {
			if newIssue.Title == openIssue.Title && openIssue.State == "opened" {
				issueExists = true
				break
			}
		}
		if !issueExists {
			reporterPlugin.CreateIssue(httpClient, strconv.Itoa(project.Id), authToken, newIssue)
			pterm.Info.Println("new issue opened for " + newIssue.Title)
		}
	}
}

// scan container registries and collect cves
// return them afterwards
func containerRegistryScan(httpClient types.HttpClient, apiPlugin plugins.Api, project types.RasicProject, userName string, authToken string, newIssues []types.RasicIssue) []types.RasicIssue {
	// look for container registries in the project
	containerRegistries := apiPlugin.GetRepositories(httpClient, strconv.Itoa(project.Id), authToken)

	// scan all registries
	// exclude /cache ones
	// append all cves to newIssues
	for _, reg := range containerRegistries {
		containerRegistry := apiPlugin.GetRepository(httpClient, strconv.Itoa(reg.Id), authToken)

		// skip cache registires
		// only valid for kaniko projects
		if strings.Contains(containerRegistry.Tag.Location, "/cache") {
			pterm.Info.Printfln("skip cache: " + containerRegistry.Tag.Location)
			continue
		}

		pterm.Info.Printfln("scan image: " + containerRegistry.Tag.Location)
		tmpIssues, _ := scan.ContainerScanner(httpClient, apiPlugin, project, containerRegistry, authToken, userName, newIssues)
		newIssues = append(newIssues, tmpIssues...)
	}
	return newIssues
}
