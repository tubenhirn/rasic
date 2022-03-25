package main

import (
	"os"
	"time"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"tubenhirn.com/risc/commands"
)

func main() {
	app := &cli.App{
		Name:                 "risc",
		HelpName:             "",
		Usage:                "create issues from known cve's",
		UsageText:            "",
		ArgsUsage:            "",
		Version:              "v0.1.1",
		Description:          "a simple app to create issues (on gitlab.com) from known cve's",
		Commands:             []*cli.Command{},
		Flags:                []cli.Flag{},
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
		Copyright: "(c) 2021 tubenhirn.com",
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
	app.Commands = append(app.Commands, commands.List())

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		pterm.Error.Println(err)
	}
}
