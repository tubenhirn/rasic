package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/types"
)

// TODO: put base config in a config file in local .config folder
var baseUrl = "https://gitlab.com"
var apiPath = "/api/v4/"

// do a get api call against gitlab.com
func apiCallGet(client types.HttpClient, url string, token string) (*http.Response, error) {
	req, reqerr := http.NewRequest("GET", url, nil)

	if reqerr != nil {
		return nil, cli.NewExitError(reqerr, 1)
	}

	// set auth header
	req.Header.Set("PRIVATE-TOKEN", token)

	// do the request
	res, reserr := client.Do(req)
	if reserr != nil {
		return nil, cli.NewExitError(reserr, 1)
	}

	// retrun the response
	return res, nil
}

// do a post api call against gitlab.com
func apiCallPost(client types.HttpClient, url string, token string, body string) (*http.Response, error) {
	req, reqerr := http.NewRequest("POST", url, strings.NewReader(body))

	if reqerr != nil {
		return nil, cli.NewExitError(reqerr, 1)
	}

	// set auth header
	req.Header.Set("PRIVATE-TOKEN", token)
	req.Header.Set("content-type", "application/json")

	// do the request
	res, reserr := client.Do(req)
	if reserr != nil {
		return nil, cli.NewExitError(reserr, 1)
	}

	return res, nil
}

// get a list of projects in a given group
func GetProjects(client types.HttpClient, group string, token string) (types.GitlabProjects, error) {
	url := baseUrl + apiPath + "groups/" + group + "/projects?per_page=100&include_subgroups=true&archived=false"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil, cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" {
		var projectlist types.GitlabProjects
		if err := json.NewDecoder(res.Body).Decode(&projectlist); err != nil {
			pterm.Error.Println(err)
			return nil, cli.NewExitError(err, 2)
		}
		return projectlist, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return nil, cli.NewExitError(err, 3)
		}
		return nil, cli.NewExitError("response error", 2)
	}

}

// get a project
func GetProject(client types.HttpClient, project string, token string) (types.GitlabProject, error) {
	url := baseUrl + apiPath + "projects/" + project

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return types.GitlabProject{}, cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" {
		var project types.GitlabProject
		if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
			pterm.Info.Println(res.Body)
			return types.GitlabProject{}, cli.NewExitError("decoder error", 2)
		}
		return project, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return types.GitlabProject{}, cli.NewExitError("read error", 3)
		}
		return types.GitlabProject{}, cli.NewExitError("response error", 2)
	}

}

// get a list of issues from a project
func GetIssues(client types.HttpClient, project string, token string) (types.GitlabIssues, error) {
	url := baseUrl + apiPath + "projects/" + project + "/issues?per_page=100"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil, cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" {
		var issuelist types.GitlabIssues
		if err := json.NewDecoder(res.Body).Decode(&issuelist); err != nil {
			return nil, cli.NewExitError("decoder error", 2)
		}
		return issuelist, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return nil, cli.NewExitError("read error", 3)
		}
		return nil, cli.NewExitError(string(res.Status), 2)
	}

}

// get a file from a gitlab project (raw as string)
// used to get .trivyignore when scanning projects
func GetFile(client types.HttpClient, project string, filepath string, fileref string, token string) (string, error) {
	url := baseUrl + apiPath + "projects/" + project + "/repository/files/" + filepath + "/raw?ref=" + fileref

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return "", cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" {
		fileContent, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			return "", readErr
		}
		return string(fileContent), nil
	}

	return "", errors.New("no ignorefile found in project")
}

func CreateIssue(client types.HttpClient, project string, token string, issue *types.Issue) (*types.GitlabIssue, error) {
	url := baseUrl + apiPath + "projects/" + project + "/issues"

	body, marshalErr := json.Marshal(issue)
	if marshalErr != nil {
		pterm.Error.Println(marshalErr)
		return &types.GitlabIssue{}, marshalErr
	}

	res, err := apiCallPost(client, url, token, string(body))
	if err != nil {
		pterm.Error.Println(err)
		return &types.GitlabIssue{}, cli.NewExitError(err, 1)
	}

	if res.Status == "201 Created" {
		var issue types.GitlabIssue
		if err := json.NewDecoder(res.Body).Decode(&issue); err != nil {
			return &types.GitlabIssue{}, cli.NewExitError("decoder error", 2)
		}
		return &issue, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return &types.GitlabIssue{}, cli.NewExitError("read error", 3)
		}
		return &types.GitlabIssue{}, cli.NewExitError(string(res.Status), 2)
	}
}
