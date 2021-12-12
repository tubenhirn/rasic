package commands

import (
	"encoding/json"
	"fmt"
	"os"
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
			&cli.StringFlag{Name: "quiet", Usage: "silence std.out", Required: false},
		},
		Action: func(c *cli.Context) error {
			quiet := c.Bool("quiet")
			if !quiet {
				fmt.Println("scan for cve's")
			}

			app := "trivy"
			arg0 := "-q"
			arg1 := "fs"
			arg2 := "--format=json"
			arg3, _ := os.Getwd()
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
						issue.Open(&cve)
					}
				}
			}
			return nil
		},
	}

}
