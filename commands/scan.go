package commands

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"gitlab.com/jstang/rasic/core"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

var (
	registryExcludeFlage = cli.StringFlag{
		Name:        "registryexclude",
		Aliases:     []string{},
		Usage:       "exclude string to match against registry path",
		EnvVars:     []string{},
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		TakesFile:   false,
		Value:       "/cache",
		DefaultText: "",
		Destination: new(string),
		HasBeenSet:  false,
	}
	severityFlag = cli.StringFlag{
		Name:        "severity",
		Aliases:     []string{},
		Usage:       "the minimum severity a cve needs to be reported",
		EnvVars:     []string{"RASIC_SEVERITY"},
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		TakesFile:   false,
		Value:       "CRITICAL",
		DefaultText: "",
		Destination: new(string),
		HasBeenSet:  false,
	}
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
		Usage:       "a oauth token for the source provider",
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
		Usage:       "a username used by trivy container-image scanning",
		EnvVars:     []string{"RASIC_USERNAME"},
		FilePath:    "",
		Required:    false,
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
		EnvVars:     []string{"RASIC_IGNOREFILE"},
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
			projectID := c.String("project")
			userName := c.String("user")
			authToken := c.String("token")
			ignoreFileName := c.String("ignorefile")
			registryExclude := c.String("registryexclude")
			var severity types.Severity
			severity = severity.FromString(strings.ToTitle(c.String("severity")))

			scanContainers := c.Bool("container")

			var sourcehandshakeConfig = plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "SOURCE_PLUGIN",
				MagicCookieValue: "allow",
			}
			var reporterhandshakeConfig = plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "REPORTER_PLUGIN",
				MagicCookieValue: "allow",
			}

			var sourcePluginMap = map[string]plugin.Plugin{
				"gitlab": &plugins.SourcePlugin{},
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
					PluginPath:   "source",
					PluginHome:   pluginHome,
					PluginName:   sourceName,
					PluginConfig: sourcehandshakeConfig,
					PluginMap:    sourcePluginMap,
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
			apiPlugin, reporterPlugin, clients := core.DispensePlugins(pluginData, logger)

			for _, pluginClient := range clients {
				defer pluginClient.Kill()
			}

			pterm.Info.Println("scan for cve's")

			projects := apiPlugin.GetProjects(httpClient, projectID, authToken)
			if len(projects) < 1 {
				pterm.Info.Println("no projects found in group " + projectID + "(maybe it is a project?)")

				singleProject := apiPlugin.GetProject(httpClient, projectID, authToken)
				var currentProject types.RasicProject
				currentProject.ID = singleProject.ID
				currentProject.WebURL = singleProject.WebURL
				currentProject.DefaultBranch = singleProject.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				var newIssues []types.RasicIssue

				// scan current projects repositry (fs)
				pterm.Info.Printfln("scan repository: " + currentProject.WebURL)
				tmpIssues, err := core.RepositoryScanner(httpClient, apiPlugin, currentProject, authToken, newIssues, severity)
				newIssues = append(newIssues, tmpIssues...)
				if err != nil {
					pterm.Error.Println(err)
				}

				// scan the project contaienr registry if enabled
				if scanContainers == true {
					newIssues = core.ContainerRegistryScan(httpClient, apiPlugin, currentProject, userName, authToken, newIssues, severity, registryExclude)
				}

				newIssueCount := len(newIssues)
				if newIssueCount > 0 {
					pterm.Warning.Println(strconv.Itoa(newIssueCount) + " " + severity.String() + " issues found, check existing...")
					core.CheckLabels(httpClient, reporterPlugin, currentProject, authToken)
					core.OpenNewIssues(httpClient, reporterPlugin, currentProject, newIssues, authToken)
				} else {
					pterm.Info.Println("no issues found")
				}
				return nil
			}

			// scan a group
			pterm.Info.Println(strconv.Itoa(len(projects)) + " projects found in group " + projectID)
			for _, project := range projects {
				var newIssues []types.RasicIssue

				var currentProject types.RasicProject
				currentProject.ID = project.ID
				currentProject.WebURL = project.WebURL
				currentProject.DefaultBranch = project.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				pterm.Info.Println("scan: " + project.WebURL)

				tmpIssues, err := core.RepositoryScanner(httpClient, apiPlugin, currentProject, authToken, newIssues, severity)
				newIssues = append(newIssues, tmpIssues...)
				if err != nil {
					pterm.Error.Println(err)
				}

				// scan the project contaienr registry if enabled
				if scanContainers == true {
					newIssues = core.ContainerRegistryScan(httpClient, apiPlugin, currentProject, userName, authToken, newIssues, severity, registryExclude)
				}

				newIssueCount := len(newIssues)
				if newIssueCount > 0 {
					pterm.Warning.Println(strconv.Itoa(newIssueCount) + " " + severity.String() + " issues found, check existing...")
					core.CheckLabels(httpClient, reporterPlugin, currentProject, authToken)
					core.OpenNewIssues(httpClient, reporterPlugin, currentProject, newIssues, authToken)
				} else {
					pterm.Info.Println("no issues found")
				}
			}

			return nil
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Subcommands:            []*cli.Command{},
		Flags:                  []cli.Flag{&projectFlag, &tokenFlag, &userNameFlag, &ignoreFileFlag, &containerScannerFlag, &severityFlag, &registryExcludeFlage},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
