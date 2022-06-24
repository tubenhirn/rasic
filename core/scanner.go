package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
	"golang.org/x/exp/slices"

	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

// create a local files for additional configuration for projects
// downloaded from the respective project given
func createLocalTempfile(client types.HTTPClient, source plugins.Source, projectID string, fileName string, defaultBranch string, authToken string) (string, error) {
	fileString := source.GetFile(client, projectID, fileName, defaultBranch, authToken)
	filePath := "/tmp/scan-" + projectID + "/"

	if len(fileString) > 0 {
		pterm.Info.Println("found " + fileName + " file in project")
		dirErr := os.Mkdir(filePath, 0755)
		if dirErr != nil {
			pterm.Warning.Println(dirErr)
		}
		file, fileCreateError := os.Create(filePath + fileName)
		if fileCreateError != nil {
			return "", fileCreateError
		}
		_, err := file.WriteString(fileString)
		if err != nil {
			file.Close()
			return "", err
		}
		err = file.Close()
		if err != nil {
			return "", err
		}
		return (filePath + fileName), nil
	}
	pterm.Info.Println("no " + fileName + " file in project - skip")
	return "", nil
}

// remove temp dir used for project ignorefile
func cleanTempFiles(fileName string) {
	tempDir, _ := path.Split(fileName)
	os.RemoveAll(tempDir)
}

// scan a remote repository
func RepositoryScanner(client types.HTTPClient, source plugins.Source, project types.RasicProject, token string, knownIssues []types.RasicIssue, minSeverity types.Severity) ([]types.RasicIssue, error) {
	// find path to trivy binary
	binary, lookErr := exec.LookPath("trivy")
	if lookErr != nil {
		pterm.Error.Println(lookErr)
	}

	resultfilePath := "/tmp/scan-" + strconv.Itoa(project.ID) + "/repo_result.json"

	// build args for repo scanning
	commandArgs := []string{"-q", "repo", "--format=json", "--output=" + resultfilePath}

	// look for a ignorefile in the project
	// if it exists download it
	if project.IgnoreFileName != "" {
		ignorefilePath, _ := createLocalTempfile(client, source, strconv.Itoa(project.ID), project.IgnoreFileName, project.DefaultBranch, token)
		ignorefileArg := "--ignorefile=" + ignorefilePath
		commandArgs = append(commandArgs, ignorefileArg)

		defer cleanTempFiles(ignorefilePath)
	}

	// append project to args
	commandArgs = append(commandArgs, project.WebURL)

	// set auth var for trivy - following the docs for scanning a remote repositry
	// https://aquasecurity.github.io/trivy/v0.25.0/vulnerability/scanning/git-repository/
	os.Setenv("GITLAB_TOKEN", token)
	// unset private token. only one token can be set
	// TODO: check uri for github or gitlab and
	// decide what token to unset
	os.Unsetenv("GITHUB_TOKEN")

	// get current environment
	env := os.Environ()

	// exec trivy with commandArgs and env
	cmd := exec.Command(binary, commandArgs...)

	// use current env for execution
	cmd.Env = env

	_, execErr := cmd.Output()
	if execErr != nil {
		return nil, execErr
	}

	repoResult, err := ioutil.ReadFile(resultfilePath)
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
func containerScanner(client types.HTTPClient, source plugins.Source, project types.RasicProject, repository types.RasicRepository, token string, user string, knownIssues []types.RasicIssue, minSeverity types.Severity) ([]types.RasicIssue, error) {
	// find path to trivy binary
	binary, lookErr := exec.LookPath("trivy")
	if lookErr != nil {
		pterm.Error.Println(lookErr)
	}

	resultfilePath := "/tmp/scan-" + strconv.Itoa(project.ID) + "/image_result.json"

	// build args for image scanning
	commandArgs := []string{"-q", "image", "--format=json", "--output=" + resultfilePath}

	// look for a ignorefile in the project
	// if it exists download it
	if project.IgnoreFileName != "" {
		ignorefilePath, _ := createLocalTempfile(client, source, strconv.Itoa(project.ID), project.IgnoreFileName, project.DefaultBranch, token)
		ignorefileArg := "--ignorefile=" + ignorefilePath
		commandArgs = append(commandArgs, ignorefileArg)

		defer cleanTempFiles(ignorefilePath)
	}

	// append project to args
	commandArgs = append(commandArgs, repository.Tag.Location)

	// set auth vars for trivy - following the docs for scanning a private container registry
	// https://aquasecurity.github.io/trivy/v0.25.4/docs/advanced/private-registries/docker-hub/
	if token != "" {
		os.Setenv("TRIVY_PASSWORD", token)
	}
	// set an empty username if give
	// this is required for trivy repo scanning with gcr TRIVY_USERNAME=""
	// https://aquasecurity.github.io/trivy/v0.27.1/docs/advanced/private-registries/gcr/
	if user != "" {
		os.Setenv("TRIVY_USERNAME", user)
	}
	// get current environment
	env := os.Environ()

	// exec trivy with commandArgs and env
	cmd := exec.Command(binary, commandArgs...)

	// use current env for execution
	cmd.Env = env

	_, execErr := cmd.Output()
	if execErr != nil {
		return nil, execErr
	}

	repoResult, err := ioutil.ReadFile(resultfilePath)
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
// this need to be done if a porject contains multiple images
// or if the fs scan and the image scan have found similar cve's
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
						newIssue, _ := Template(strconv.Itoa(project.ID), cve, result.Target, result.Type)
						// add newIssue to the result list
						issueList = append(issueList, newIssue)
						// add issue title to the cveSlice for dublication checking
						cveSlice = append(cveSlice, newIssue.Title)
					}
				}
			}
		}
	}
	return issueList
}

// scan container registries and collect cves
// return them afterwards
func ContainerRegistryScan(httpClient types.HTTPClient, apiPlugin plugins.Source, project types.RasicProject, userName string, authToken string, newIssues []types.RasicIssue, severity types.Severity, registryExcudePattern string) []types.RasicIssue {
	// look for a rasic config file in the project
	// if it exists download it
	configfileName := ".rasicrc"
	configfilePath, _ := createLocalTempfile(httpClient, apiPlugin, strconv.Itoa(project.ID), configfileName, project.DefaultBranch, authToken)

	var projectConfiguration types.RasicConfiguration
	if configfilePath != "" {
		projectConfiguration, _ = buildRepositoryConfiguration(configfilePath)
	}

	// if a custom registry is configured (e.g. gcr)
	// use it instead of looking for attached ones
	if projectConfiguration.Repository.Tag.Location != "" {
		pterm.Info.Printfln("scan image: " + projectConfiguration.Repository.Tag.Location)
		tmpIssues, _ := containerScanner(httpClient, apiPlugin, project, projectConfiguration.Repository, "", "", newIssues, severity)
		newIssues = append(newIssues, tmpIssues...)
	} else {
		// look for container registries attached to the project (gitlab)
		containerRegistries := apiPlugin.GetRepositories(httpClient, strconv.Itoa(project.ID), authToken)

		// scan all registries
		// exclude /cache ones
		// append all cves to newIssues
		for _, reg := range containerRegistries {
			containerRegistry := apiPlugin.GetRepository(httpClient, strconv.Itoa(reg.ID), authToken)

			// skip cache registires
			if strings.Contains(containerRegistry.Tag.Location, registryExcudePattern) {
				pterm.Info.Printfln("skip registry: " + containerRegistry.Tag.Location + " found " + registryExcudePattern)
				continue
			}

			pterm.Info.Printfln("scan image: " + containerRegistry.Tag.Location)
			tmpIssues, _ := containerScanner(httpClient, apiPlugin, project, containerRegistry, authToken, userName, newIssues, severity)
			newIssues = append(newIssues, tmpIssues...)
		}
	}
	return newIssues
}

// build a RasicConfiguration
// return a config struct
func buildRepositoryConfiguration(configFilePath string) (types.RasicConfiguration, error) {
	var config types.RasicConfiguration

	// open config file from given path
	configFile, err := os.Open(configFilePath)
	if err != nil {
		pterm.Error.Println(err)
		return types.RasicConfiguration{}, err
	}

	// decode the jsonfile to a
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		pterm.Error.Println(err)
		return types.RasicConfiguration{}, err
	}

	return config, nil
}
