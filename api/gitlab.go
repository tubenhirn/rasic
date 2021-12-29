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

func GetProjectList(group string, token string) (types.Projects, error) {
	url := baseUrl + apiPath + "groups/" + group + "/projects?per_page=100&include_subgroups=true&archived=false"
	client := http.Client{}
	req, reqerr := http.NewRequest("GET", url, nil)
	if reqerr != nil {
		pterm.Error.Println(reqerr)
		return nil, cli.NewExitError("request error", 1)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, reserr := client.Do(req)
	if reserr != nil {
		pterm.Error.Println(reserr)
		return nil, cli.NewExitError("request error", 1)
	}

	defer res.Body.Close()

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

func GetProject(project string, token string) (*types.Project, error) {
	url := baseUrl + apiPath + "projects/" + project
	client := http.Client{}
	req, reqerr := http.NewRequest("GET", url, nil)
	if reqerr != nil {
		pterm.Error.Println(reqerr)
		return nil, cli.NewExitError("request error", 1)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, reserr := client.Do(req)
	if reserr != nil {
		pterm.Error.Println(reserr)
		return nil, cli.NewExitError("request error", 1)
	}

	defer res.Body.Close()

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

func GetIssueList(project string, token string) (types.Issues, error) {
	url := baseUrl + apiPath + "projects/" + project + "/issues?per_page=100"
	client := http.Client{}
	req, reqerr := http.NewRequest("GET", url, nil)
	if reqerr != nil {
		pterm.Error.Println(reqerr)
		return nil, cli.NewExitError("request error", 1)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, reserr := client.Do(req)
	if reserr != nil {
		pterm.Error.Println(reserr)
		return nil, cli.NewExitError("request error", 1)
	}

	defer res.Body.Close()

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
func GetFile(project string, filepath string, fileref string, token string) (string, error) {
	url := baseUrl + apiPath + "projects/" + project + "/repository/files/" + filepath + "/raw?ref=" + fileref

	client := http.Client{}
	req, reqerr := http.NewRequest("GET", url, nil)
	if reqerr != nil {
		pterm.Error.Println(reqerr)
		return "", cli.NewExitError("request error", 1)
	}

	req.Header.Set("PRIVATE-TOKEN", token)

	res, reserr := client.Do(req)
	if reserr != nil {
		pterm.Error.Println(reserr)
		return "", cli.NewExitError("request error", 1)
	}
	if res.Status == "200 OK" || res.Status == "200" {
		fileContent, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			return "", readErr
		}
		return string(fileContent), nil
	}

	defer res.Body.Close()

	return "", errors.New("no ignorefile found in project")
}
