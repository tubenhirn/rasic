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
func createLocalIgnorefile(client types.HttpClient, source plugins.Api, projectId string, ignoreFileName string, defaultBranch string, authToken string) (string, error) {
	ignorefileString := source.GetFile(client, projectId, ignoreFileName, defaultBranch, authToken)
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

func RepositoryScanner(client types.HttpClient, source plugins.Api, project types.RasicProject, token string, knownIssues []types.RasicIssue) ([]types.RasicIssue, error) {

	// look for a ignorefile in the project
	// if it exists download it
	ignorefilePath, _ := createLocalIgnorefile(client, source, strconv.Itoa(project.Id), project.IgnoreFileName, project.DefaultBranch, token)
	defer cleanTempFiles(ignorefilePath)

	// find path to trivy binary
	binary, lookErr := exec.LookPath("trivy")
	if lookErr != nil {
		pterm.Error.Println(lookErr)
	}
	// build args for repo scanning
	repoArgs := []string{"-q", "repo", "--ignorefile=" + ignorefilePath, "--format=json", "--output=repo_result.json", project.WebUrl}

	// build args for image scanning
	// imageArgs := []string{"-q", "image", "--ignorefile=" + ignorefilePath, "--format=json", "--output=image_result.json", project.WebUrl}

	// set auth var for trivy - following the docs for scanning a remote repositry
	// https://aquasecurity.github.io/trivy/v0.25.0/vulnerability/scanning/git-repository/
	os.Setenv("GITLAB_TOKEN", token)

	// get current environment
	env := os.Environ()

	// exec trivy with repoArgs and env
	cmd := exec.Command(binary, repoArgs...)

	// use current env for execution
	cmd.Env = env

	_, execErr := cmd.Output()
	if execErr != nil {
		return nil, execErr
	}

	repoResult, err := ioutil.ReadFile("repo_result.json")
	if err != nil {
		pterm.Error.Printf("Status: %s\n", "file read error")
	}

	var report types.CVEReport
	unmarshalerr := json.Unmarshal(repoResult, &report)
	if unmarshalerr != nil {
		pterm.Error.Println(unmarshalerr)
		return nil, unmarshalerr
	}

	var issueList []types.RasicIssue
	for _, result := range report.Results {
		if len(result.Vulnerabilities) > 0 {

			pterm.Warning.Println(strconv.Itoa(len(result.Vulnerabilities)) + " " + result.Type + " vulnerabilities found")

			for _, cve := range result.Vulnerabilities {
				// check for open knownIssues
				exists := false
				for i := range knownIssues {
					if knownIssues[i].Title == cve.VulnerabilityID {
						// TODO: allow to control State
						// maybe check for label "wont-fix" on the closed issue - otherwise reopen it
						if knownIssues[i].State == "opened" {
							pterm.Info.Println("open issue exists for " + cve.VulnerabilityID)
							exists = true
							break
						}
					}
				}
				if !exists {
					// create new issue and add it to the list we return
					newIssue, _ := issue.Template(strconv.Itoa(project.Id), cve, result.Target, result.Type)

					// TODO: make this configurable
					// and better.....
					minSeverity := "CRITICAL"
					if cve.Severity == minSeverity {
						issueList = append(issueList, newIssue)
					}
				}

			}
		}
	}
	return issueList, nil
}

func ContainerScanner(client types.HttpClient, source plugins.Api, project types.RasicProject, repository types.RasicRepository, token string, user string, knownIssues []types.RasicIssue) ([]types.RasicIssue, error) {

	// look for a ignorefile in the project
	// if it exists download it
	ignorefilePath, _ := createLocalIgnorefile(client, source, strconv.Itoa(project.Id), project.IgnoreFileName, project.DefaultBranch, token)
	defer cleanTempFiles(ignorefilePath)

	// find path to trivy binary
	binary, lookErr := exec.LookPath("trivy")
	if lookErr != nil {
		pterm.Error.Println(lookErr)
	}

	// build args for image scanning
	imageArgs := []string{"-q", "image", "--ignorefile=" + ignorefilePath, "--format=json", "--output=image_result.json", repository.Tag.Location}

	// set auth vars for trivy - following the docs for scanning a private container registry
	// https://aquasecurity.github.io/trivy/v0.25.4/docs/advanced/private-registries/docker-hub/
	os.Setenv("TRIVY_PASSWORD", token)
	os.Setenv("TRIVY_USERNAME", user)

	// get current environment
	env := os.Environ()

	// exec trivy with repoArgs and env
	cmd := exec.Command(binary, imageArgs...)

	// use current env for execution
	cmd.Env = env

	_, execErr := cmd.Output()
	if execErr != nil {
		return nil, execErr
	}

	repoResult, err := ioutil.ReadFile("image_result.json")
	if err != nil {
		pterm.Error.Printf("Status: %s\n", "file read error")
	}

	var report types.CVEReport
	unmarshalerr := json.Unmarshal(repoResult, &report)
	if unmarshalerr != nil {
		pterm.Error.Println(unmarshalerr)
		return nil, unmarshalerr
	}

	// build a list of isses
	// check for known ones to dont add them twice
	// this need to be done if a porject containts multiple images
	// or if the fs scan and the image scan have found similiar cve's
	// maybe this can be removed in the future
	// we also only add cve's with a give severity
	var issueList []types.RasicIssue
	for _, result := range report.Results {
		if len(result.Vulnerabilities) > 0 {

			pterm.Warning.Println(strconv.Itoa(len(result.Vulnerabilities)) + " " + result.Type + " vulnerabilities found")

			for _, cve := range result.Vulnerabilities {
				exists := false
				for i := range knownIssues {
					if knownIssues[i].Title == cve.VulnerabilityID {
						pterm.Info.Println("open issue exists for " + cve.VulnerabilityID)
						exists = true
						break
					}
				}
				if !exists {
					// TODO: make this configurable
					// and better.....
					minSeverity := "CRITICAL"
					if cve.Severity == minSeverity {
						// create new issue and add it to the list we return
						newIssue, _ := issue.Template(strconv.Itoa(project.Id), cve, result.Target, result.Type)
						issueList = append(issueList, newIssue)
					}
				}

			}
		}
	}
	return issueList, nil
}
