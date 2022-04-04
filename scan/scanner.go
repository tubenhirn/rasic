package scan

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/pterm/pterm"

	"gitlab.com/jstang/rasic/issue"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

// create a local .triviyignore file
// downloaded from the respective project given
func createLocalIgnorefile(client types.HttpClient, api plugins.Api, projectId string, ignoreFileName string, defaultBranch string, authToken string) (string, error) {
	ignorefileString := api.GetFile(client, projectId, ignoreFileName, defaultBranch, authToken)
	ignoreFilePath := "/tmp/scan-" + projectId + "/"

	if len(ignorefileString) > 0 {
		pterm.Info.Println("found " + ignoreFileName + "file in project")
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

/**
  scan with trivy binary
  and save the output as result.json
 **/

func Scanner(client types.HttpClient, api plugins.Api, project types.RasicProject, token string, issues []types.RasicIssue) error {

	// look for a ignorefile in the project
	// if it exists download it
	ignorefilePath, _ := createLocalIgnorefile(client, api, strconv.Itoa(project.Id), project.IgnoreFileName, project.DefaultBranch, token)
	defer cleanTempFiles(ignorefilePath)

	// find path to trivy binary
	binary, lookErr := exec.LookPath("trivy")
	if lookErr != nil {
		panic(lookErr)
	}
	// build args
	args := []string{"-q", "repo", "--ignorefile=" + ignorefilePath, "--format=json", "--output=result.json", project.WebUrl}

	// set auth var for trivy - following the docs for scanning a remote repositry
	// https://aquasecurity.github.io/trivy/v0.25.0/vulnerability/scanning/git-repository/
	os.Setenv("GITLAB_TOKEN", token)

	// get current environment
	env := os.Environ()

	// exec trivy with args and env
	cmd := exec.Command(binary, args...)

	// use current env for execution
	cmd.Env = env

	_, execErr := cmd.Output()
	if execErr != nil {
		pterm.Error.Printfln(execErr.Error())
		return execErr
	}

	result, err := ioutil.ReadFile("result.json")
	if err != nil {
		pterm.Error.Printf("Status: %s\n", "file read error")
	}

	var report types.CVEReport
	unmarshalerr := json.Unmarshal(result, &report)
	if unmarshalerr != nil {
		pterm.Error.Println(unmarshalerr)
		return unmarshalerr
	}

	for _, result := range report.Results {
		if len(result.Vulnerabilities) > 0 {

			pterm.Warning.Println(strconv.Itoa(len(result.Vulnerabilities)) + " " + result.Type + " vulnerabilities found")

			for _, cve := range result.Vulnerabilities {
				// check for open issues
				exists := false
				for i := range issues {
					if issues[i].Title == cve.VulnerabilityID {
						// TODO: allow to control State
						// maybe check for label "wont-fix" on the closed issue - otherwise reopen it
						if issues[i].State == "opened" {
							pterm.Info.Println("open issue exists for " + cve.VulnerabilityID)
							exists = true
							break
						}
					}
				}
				if !exists {
					// open issue if no issuw present in thes current porject
					newIssue, _ := issue.Template(strconv.Itoa(project.Id), cve, result.Target, result.Type)

					// TODO: make this configurable
					// and better.....
					if cve.Severity == "HIGH" {
						api.CreateIssue(client, strconv.Itoa(project.Id), token, newIssue)
						pterm.Info.Println("new issue opened for " + cve.VulnerabilityID + " - " + cve.Severity)
					}
				}

			}
		}
	}
	return nil
}
