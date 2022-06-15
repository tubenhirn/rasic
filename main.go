package main

import (
	"encoding/gob"
	"net/http"
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"gitlab.com/jstang/rasic/commands"
	"gitlab.com/jstang/rasic/types"
)

// passed with ldflag
var appVersion string

// register types to gob
// this is required to proper serialize and deserialize the data
func init() {
	gob.Register(http.DefaultClient)
	gob.Register(types.RasicIssue{})
	gob.Register(types.RasicLabel{})
	gob.Register(types.RasicProject{})
	gob.Register(types.RasicRepository{})
	gob.Register(map[string]interface{}{})
}

func main() {
	// determine current user dir
	// this is required for the pluginhome directory
	userHome, homeErr := os.UserHomeDir()
	if homeErr != nil {
		pterm.Error.Println(homeErr)
	}

	app := &cli.App{
		Name:        "rasic",
		HelpName:    "",
		Usage:       "create issues from known cve's",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     appVersion,
		Description: "a simple app to create issues for known cve's or config flaws",
		Commands:    []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "source",
				Aliases:     []string{},
				Usage:       "specify the source of your files (gitlab, github)",
				EnvVars:     []string{"RASIC_SOURCE"},
				FilePath:    "",
				Required:    false,
				Hidden:      false,
				TakesFile:   false,
				Value:       "gitlab",
				DefaultText: "",
				Destination: new(string),
				HasBeenSet:  false,
			},
			&cli.StringFlag{
				Name:        "reporter",
				Aliases:     []string{},
				Usage:       "specify the platform to create issues in (gitlab, github, jira)",
				EnvVars:     []string{"RASIC_REPORTER"},
				FilePath:    "",
				Required:    false,
				Hidden:      false,
				TakesFile:   false,
				Value:       "gitlab",
				DefaultText: "",
				Destination: new(string),
				HasBeenSet:  false,
			},
			&cli.StringFlag{
				Name:        "scanner",
				Aliases:     []string{},
				Usage:       "specify the scanner tool (trivy, tfsec...)",
				EnvVars:     []string{"RASIC_SCANNER"},
				FilePath:    "",
				Required:    false,
				Hidden:      false,
				TakesFile:   false,
				Value:       "trivy",
				DefaultText: "",
				Destination: new(string),
				HasBeenSet:  false,
			},
			&cli.StringFlag{
				Name:        "pluginhome",
				Aliases:     []string{},
				Usage:       "specify the location your plugins are stored",
				EnvVars:     []string{"RASIC_PLUGINHOME"},
				FilePath:    "",
				Required:    false,
				Hidden:      false,
				TakesFile:   false,
				Value:       userHome + "/.rasic/plugins/",
				DefaultText: "",
				Destination: new(string),
				HasBeenSet:  false,
			},
		},
		EnableBashCompletion: false,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete: func(*cli.Context) {
		},
		Before: func(*cli.Context) error {
			return nil
		},
		After: func(*cli.Context) error {
			return nil
		},
		Action: func(*cli.Context) error {
			return nil
		},
		CommandNotFound: func(*cli.Context, string) {
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Compiled:  time.Now(),
		Authors:   []*cli.Author{{Name: "Johannes Stang", Email: "tubenhirn@gmail.com"}},
		Copyright: "(c) 2022 tubenhirn.com",
		Reader:    nil,
		Writer:    nil,
		ErrWriter: nil,
		ExitErrHandler: func(context *cli.Context, err error) {
		},
		Metadata: map[string]interface{}{},
		ExtraInfo: func() map[string]string {
			return nil
		},
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}

	app.Commands = append(app.Commands, commands.Scan())
	app.Commands = append(app.Commands, commands.Projects())
	app.Commands = append(app.Commands, commands.Issues())

	app.EnableBashCompletion = true

	app.Action = func(c *cli.Context) error {
		err := cli.ShowAppHelp(c)
		if err != nil {
			pterm.Error.Println(err)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		pterm.Error.Println(err)
	}
}
