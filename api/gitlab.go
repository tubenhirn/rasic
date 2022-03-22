package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"tubenhirn.com/cve2issue/types"
)

var baseUrl = "https://gitlab.com"
var apiPath = "/api/v4/"

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func apiCall(client HttpClient, url string, token string) (*http.Response, error) {
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
	// close the request on function end
	defer res.Body.Close()

	// retrun the response
	return res, nil
}

func GetProjectList(client *http.Client, group string, token string) (types.Projects, error) {
	url := baseUrl + apiPath + "groups/" + group + "/projects?per_page=100&include_subgroups=true&archived=false"

	res, err := apiCall(client, url, token)
	if err != nil {
		pterm.Error.Println(err)
		return nil, cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" || res.Status == "200" {
		var projectlist types.Projects
		if err := json.NewDecoder(res.Body).Decode(&projectlist); err != nil {
			return nil, cli.NewExitError("decoder error", 2)
		}
		return projectlist, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return nil, cli.NewExitError("read error", 3)
		}
		return nil, cli.NewExitError(string(res.Status), 2)
	}

}

func GetProject(client *http.Client, project string, token string) (*types.Project, error) {
	url := baseUrl + apiPath + "projects/" + project

	res, err := apiCall(client, url, token)
	if err != nil {
		pterm.Error.Println(err)
		return nil, cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" || res.Status == "200" {
		project := &types.Project{}
		if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
			return nil, cli.NewExitError("decoder error", 2)
		}
		return project, nil
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return nil, cli.NewExitError("read error", 3)
		}
		return nil, cli.NewExitError(string(res.Status), 2)
	}

}

func GetIssueList(client *http.Client, project string, token string) (types.Issues, error) {
	url := baseUrl + apiPath + "projects/" + project + "/issues?per_page=100"

	res, err := apiCall(client, url, token)
	if err != nil {
		pterm.Error.Println(err)
		return nil, cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" || res.Status == "200" {
		var issuelist types.Issues
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
func GetFile(client *http.Client, project string, filepath string, fileref string, token string) (string, error) {
	url := baseUrl + apiPath + "projects/" + project + "/repository/files/" + filepath + "/raw?ref=" + fileref

	res, err := apiCall(client, url, token)
	if err != nil {
		pterm.Error.Println(err)
		return "", cli.NewExitError(err, 1)
	}

	if res.Status == "200 OK" || res.Status == "200" {
		fileContent, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			return "", readErr
		}
		return string(fileContent), nil
	}

	return "", errors.New("no ignorefile found in project")
}
