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
		Name:    "projects",
		Aliases: []string{"p"},
		Usage:   "list projects of a group",
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
	}

}
