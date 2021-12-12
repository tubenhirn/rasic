package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"

	"tubenhirn.com/cve2issue/commands"
)

func main() {
	app := &cli.App{
		Name:        "cve2issue",
		Usage:       "create issues from known cve's",
		Description: "a simple app to create issues (on gitlab.com or jira) from known cve's",
		Version:     "v0.1.0",
		Compiled:    time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Johannes Stang",
				Email: "tubenhirn@gmail.com",
			},
		},
		Copyright: "(c) 2021 tubenhirn.com",
	}

	app.Commands = append(app.Commands, commands.Scan())

	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
