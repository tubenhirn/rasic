package commands

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"

	"tubenhirn.com/cve2issue/issue"
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
				Aliases:     []string{},
				Usage:       "a project url to scan",
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
			project := c.String("project")
			fmt.Println("scan for cve's")

			app := "trivy"
			arg0 := "-q"
			arg1 := "repo"
			arg2 := "--format=json"
			arg3 := project
			cmd := exec.Command(app, arg0, arg1, arg2, arg3)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			var report types.CVEReport
			unmarshalerr := json.Unmarshal(stdout, &report)
			if unmarshalerr != nil {
				return unmarshalerr
			}
			for _, result := range report.Results {
				if len(result.Vulnerabilities) > 0 {
					fmt.Println(result.Target)
					for _, cve := range result.Vulnerabilities {
						issue.Open("", &cve, result.Target, result.Type)
					}
				}
			}
			return nil
		},
	}

}
