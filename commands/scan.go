package commands

import (
	"strconv"

	"github.com/pterm/pterm"
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
			projectId := c.String("project")
			authToken := c.String("token")
			pterm.Info.Println("scan for cve's")
			var projects types.Projects
			projects, _ = api.GetProjectList(projectId, authToken)
			if len(projects) < 1 {
				pterm.Info.Println("no projects found in group " + projectId + "(maybe it is a project?)")
				// try to get the project
				singleProject, _ := api.GetProject(projectId, authToken)
				var issues types.Issues
				issues, _ = api.GetIssueList(strconv.Itoa(singleProject.ID), authToken)
				pterm.Info.Printfln("scan: " + singleProject.WebURL)
				err := scan.Scanner(singleProject.WebURL, issues)
				if err != nil {
					pterm.Error.Println(err)
				}

				return nil
			}

			for _, project := range projects {
				// get all issues for current project
				var issues types.Issues
				issues, _ = api.GetIssueList(strconv.Itoa(project.ID), authToken)
				pterm.Info.Println("scan: " + project.WebURL)
				err := scan.Scanner(project.WebURL, issues)
				if err != nil {
					pterm.Error.Println(err)
				}
			}

			return nil
		},
	}

}
