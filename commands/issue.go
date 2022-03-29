package commands

import (
	"encoding/json"
	"net/http"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/api"
)

// open an new issue for a found cve
func ListIssues() *cli.Command {
	return &cli.Command{
		Name:        "issues",
		Aliases:     []string{"i"},
		Usage:       "interact with issues of a project",
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
			// show default help since this is only a group command
			// see subcommands for details
			cli.ShowAppHelp(c)
			return nil
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:        "list",
				Aliases:     []string{},
				Usage:       "",
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
					project := c.String("project")
					token := c.String("token")
					client := &http.Client{}
					issues, _ := api.GetIssues(client, project, token)
					bytes, marshalerror := json.Marshal(issues)
					if marshalerror != nil {
						pterm.Error.Println(marshalerror)
					}
					pterm.Info.Println(string(bytes))

					return nil
				},
				OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
					return nil
				},
				Subcommands: []*cli.Command{},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "project",
						Aliases:     []string{},
						Usage:       "a project id",
						EnvVars:     []string{},
						FilePath:    "",
						Required:    true,
						Hidden:      false,
						TakesFile:   false,
						Value:       "",
						DefaultText: "",
						Destination: new(string),
						HasBeenSet:  false,
					},
					&cli.StringFlag{
						Name:        "token",
						Aliases:     []string{},
						Usage:       "a private token",
						EnvVars:     []string{},
						FilePath:    "",
						Required:    true,
						Hidden:      false,
						TakesFile:   false,
						Value:       "",
						DefaultText: "",
						Destination: new(string),
						HasBeenSet:  false,
					},
				},
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			},
		},
		Flags:                  []cli.Flag{},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

}
