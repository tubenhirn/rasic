package commands

import (
	"github.com/urfave/cli/v2"
)

func OpenIssue() *cli.Command {
	return &cli.Command{
		Name:    "open",
		Aliases: []string{"o"},
		Usage:   "open an issue for known cve's",
		Flags:   []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

}
