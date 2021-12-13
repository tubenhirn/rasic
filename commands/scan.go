package commands

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"

	"tubenhirn.com/cve2issue/api"
	"tubenhirn.com/cve2issue/scan"
	"tubenhirn.com/cve2issue/types"
)

func Scan() *cli.Command {
	return &cli.Command{
		Name:    "scan",
		Aliases: []string{"s"},
		Usage:   "scan project for cve's",
		Flags: []cli.Flag{
			&cli.StringFlag{
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
			},
			&cli.StringFlag{
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
			},
		},
		Action: func(c *cli.Context) error {
			project := c.String("project")
			token := c.String("token")
			fmt.Println("scan for cve's")

			var projects types.Projects
			projects, _ = api.GetProjectList(project, token)
			if len(projects) < 1 {
				fmt.Println("no projects found in group " + project + "(maybe it is project?)")
				singleProject, _ := api.GetProject(project, token)
				var issues types.Issues
				issues, _ = api.GetIssueList(strconv.Itoa(singleProject.ID), token)
				fmt.Println("scan: " + singleProject.WebURL)
				err := scan.Scanner(singleProject.WebURL, issues)
				if err != nil {
					fmt.Println(err)
				}

				return nil
			}

			for _, pro := range projects {
				// get all issues for current project
				var issues types.Issues
				issues, _ = api.GetIssueList(strconv.Itoa(pro.ID), token)
				fmt.Println("scan: " + pro.WebURL)
				err := scan.Scanner(pro.WebURL, issues)
				if err != nil {
					fmt.Println(err)
				}
			}

			return nil
		},
	}

}
