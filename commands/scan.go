package commands

import (
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"

	"tubenhirn.com/rasic/api"
	"tubenhirn.com/rasic/scan"
	"tubenhirn.com/rasic/types"
)

var (
	projectFlag = cli.StringFlag{
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
	}
	tokenFlag = cli.StringFlag{
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
	}
	ignoreFileFlag = cli.StringFlag{
		Name:        "ignorefile",
		Aliases:     []string{},
		Usage:       "specify a ignorefile. set ignorefile flag or use CVE_IGNORE_FILE env. (default .trivyignore)",
		EnvVars:     []string{"CVE_IGNORE_FILE"},
		FilePath:    "",
		Required:    false,
		Hidden:      false,
		TakesFile:   false,
		Value:       ".trivyignore",
		DefaultText: "",
		Destination: new(string),
		HasBeenSet:  false,
	}
)

// create a local .triviyignore file
// downloaded from the respective project given
func createLocalIgnorefile(client *http.Client, projectId string, ignoreFileName string, defaultBranch string, authToken string) (string, error) {
	ignorefileString, fileErr := api.GetFile(client, projectId, ignoreFileName, defaultBranch, authToken)
	ignoreFilePath := "/tmp/scan-" + projectId + "/"
	if fileErr != nil {
		return "", fileErr
	}
	if len(ignorefileString) > 0 {
		pterm.Info.Println("found .trivyignore file in project")
		dirErr := os.Mkdir(ignoreFilePath, 0755)
		if dirErr != nil {
			pterm.Warning.Println(dirErr)
		}
		file, fileCreateError := os.Create(ignoreFilePath + ignoreFileName)
		if fileCreateError != nil {
			return "", fileCreateError
		}
		_, err := file.WriteString(ignorefileString)
		if err != nil {
			file.Close()
			return "", err
		}
		err = file.Close()
		if err != nil {
			return "", err
		}
	}
	return (ignoreFilePath + ignoreFileName), nil
}

// remove temp dir used for project ignorefile
func cleanTempFiles(fileName string) error {
	tempDir, _ := path.Split(fileName)
	rmErr := os.RemoveAll(tempDir)
	if rmErr != nil {
		return rmErr
	}
	return nil
}

// scan a project for cve's
func Scan() *cli.Command {
	return &cli.Command{
		Name:        "scan",
		Aliases:     []string{"s"},
		Usage:       "scan project for cve's",
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
			projectId := c.String("project")
			authToken := c.String("token")
			ignoreFileName := c.String("ignorefile")
			pterm.Info.Println("scan for cve's")
			var projects types.Projects
			client := &http.Client{}
			projects, _ = api.GetProjectList(client, projectId, authToken)
			if len(projects) < 1 {
				pterm.Info.Println("no projects found in group " + projectId + "(maybe it is a project?)")

				singleProject, _ := api.GetProject(client, projectId, authToken)

				var issues types.Issues
				issues, _ = api.GetIssueList(client, strconv.Itoa(singleProject.ID), authToken)
				pterm.Info.Printfln("scan: " + singleProject.WebURL)

				tempFileName, _ := createLocalIgnorefile(client, strconv.Itoa(singleProject.ID), ignoreFileName, singleProject.DefaultBranch, authToken)

				err := scan.Scanner(singleProject.WebURL, issues, tempFileName)
				if err != nil {
					pterm.Error.Println(err)
				}

				defer cleanTempFiles(tempFileName)

				return nil
			}
			pterm.Info.Println(strconv.Itoa(len(projects)) + " projects found in group " + projectId)
			for _, project := range projects {
				var issues types.Issues
				issues, _ = api.GetIssueList(client, strconv.Itoa(project.ID), authToken)
				pterm.Info.Println("scan: " + project.WebURL)

				tempFileName, _ := createLocalIgnorefile(client, strconv.Itoa(project.ID), ignoreFileName, project.DefaultBranch, authToken)

				err := scan.Scanner(project.WebURL, issues, tempFileName)
				if err != nil {
					pterm.Error.Println(err)
				}
				defer cleanTempFiles(tempFileName)
			}
			return nil
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Subcommands:            []*cli.Command{},
		Flags:                  []cli.Flag{&projectFlag, &tokenFlag, &ignoreFileFlag},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

}
