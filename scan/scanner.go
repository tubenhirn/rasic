package scan

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/pterm/pterm"
	"golang.org/x/exp/slices"

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

// scan a remote repository
func RepositoryScanner(client types.HttpClient, source plugins.Api, project types.RasicProject, token string, knownIssues []types.RasicIssue, minSeverity types.Severity) ([]types.RasicIssue, error) {

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
	}
	issueList := buildIssueList(report, knownIssues, project, minSeverity)

	return issueList, nil
}

// scan containers in the project - if present
func ContainerScanner(client types.HttpClient, source plugins.Api, project types.RasicProject, repository types.RasicRepository, token string, user string, knownIssues []types.RasicIssue, minSeverity types.Severity) ([]types.RasicIssue, error) {

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
	}
	issueList := buildIssueList(report, knownIssues, project, minSeverity)

	return issueList, nil
}

// build a list of isses
// check for known ones to dont add them twice
// this need to be done if a porject containts multiple images
// or if the fs scan and the image scan have found similiar cve's
// maybe this can be removed in the future
// we also only add cve's with a give severity
func buildIssueList(report types.CVEReport, knownIssues []types.RasicIssue, project types.RasicProject, minSeverity types.Severity) []types.RasicIssue {
	var issueList []types.RasicIssue

	var cveSlice []string
	// create a list of known cves
	// used for dublication check
	for _, issue := range knownIssues {
		cveSlice = append(cveSlice, issue.Title)
	}

	// loop packages in the report
	for _, result := range report.Results {
		if len(result.Vulnerabilities) > 0 {

			pterm.Info.Println(strconv.Itoa(len(result.Vulnerabilities)) + " " + result.Type + " vulnerabilities found")

			// loop cves in the current package
			for _, cve := range result.Vulnerabilities {

				// add cve if unknown
				// and if its severity >= minSeverity
				if !slices.Contains(cveSlice, cve.VulnerabilityID) {
					var cveSeverity types.Severity
					cveSeverity = cve.Severity
					if cveSeverity >= minSeverity {
						// create new issue and add it to the list we return
						newIssue, _ := issue.Template(strconv.Itoa(project.Id), cve, result.Target, result.Type)
						issueList = append(issueList, newIssue)
					}
				}

			}
		}
	}
	return issueList
}
