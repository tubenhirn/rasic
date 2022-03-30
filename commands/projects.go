package commands

import (
	"encoding/json"
	"net/http"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"gitlab.com/jstang/rasic/api"
)

// list all projects of a give group
func ListProjects() *cli.Command {
	return &cli.Command{
		Name:        "projects",
		Aliases:     []string{"p"},
		Usage:       "interact with projects",
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
		Subcommands: []*cli.Command{{
			Name:        "list",
			Aliases:     []string{"l"},
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
				group := c.String("group")
				token := c.String("token")
				client := &http.Client{}
				projects, _ := api.GetProjects(client, group, token)
				bytes, marshalerror := json.Marshal(projects)
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
					Name:        "group",
					Aliases:     []string{},
					Usage:       "a group id",
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
		}},
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
