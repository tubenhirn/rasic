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
		Usage:       "scan project for cve's",
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
			backend := c.String("api")
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
			}

			httpClient := &http.Client{}
			logger := hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stdout,
				Level:  hclog.Error,
			})

			client := plugin.NewClient(&plugin.ClientConfig{
				HandshakeConfig: handshakeConfig,
				Plugins:         pluginMap,
				Cmd:             exec.Command("./plugins/api/" + backend),
				Logger:          logger,
			})
			defer client.Kill()

			rpcClient, err := client.Client()
			if err != nil {
				pterm.Error.Println(err)
			}

			raw, err := rpcClient.Dispense(backend)
			if err != nil {
				pterm.Error.Println(err)
			}
			api := raw.(plugins.Api)

			pterm.Info.Println("scan for cve's")
			var projects types.GitlabProjects
			projects = api.GetProjects(httpClient, projectId, authToken)
			if len(projects) < 1 {
				pterm.Info.Println("no projects found in group " + projectId + "(maybe it is a project?)")

				singleProject := api.GetProject(httpClient, projectId, authToken)
				var currentProject types.Project
				currentProject.Id = singleProject.ID
				currentProject.WebUrl = singleProject.WebURL
				currentProject.DefaultBranch = singleProject.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				var issues types.GitlabIssues
				issues = api.GetIssues(httpClient, strconv.Itoa(singleProject.ID), authToken)
				pterm.Info.Printfln("scan: " + currentProject.WebUrl)

				err = scan.Scanner(httpClient, api, currentProject, authToken, issues)
				if err != nil {
					pterm.Error.Println(err)
				}

				return nil
			}
			pterm.Info.Println(strconv.Itoa(len(projects)) + " projects found in group " + projectId)
			for _, project := range projects {
				var issues types.GitlabIssues

				var currentProject types.Project
				currentProject.Id = project.ID
				currentProject.WebUrl = project.WebURL
				currentProject.DefaultBranch = project.DefaultBranch
				currentProject.IgnoreFileName = ignoreFileName

				issues = api.GetIssues(httpClient, strconv.Itoa(project.ID), authToken)
				pterm.Info.Println("scan: " + project.WebURL)


				err := scan.Scanner(httpClient, api, currentProject, authToken, issues)
				if err != nil {
					pterm.Error.Println(err)
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
