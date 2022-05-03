package main

import (
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

var baseURL = "https://gitlab.com"
var apiPath = "/api/v4/"

const OK = "200 OK"

type ReporterGitlab struct{}

func (a *ReporterGitlab) GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject {
	url := baseURL + apiPath + "groups/" + group + "/projects?per_page=100&include_subgroups=true&archived=false"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil
	}

	if res.Status == OK {
		var projectlist types.GitlabProjects
		if err := json.NewDecoder(res.Body).Decode(&projectlist); err != nil {
			pterm.Error.Println(err)
			return nil
		}

		var returnValue []types.RasicProject
		for _, pro := range projectlist {
			ele := types.RasicProject{
				ID:            pro.ID,
				WebURL:        pro.WebURL,
				DefaultBranch: pro.DefaultBranch,
			}
			returnValue = append(returnValue, ele)
		}
		return returnValue
	}
	_, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		pterm.Error.Println(errRead)
	}
	return nil
}

func (a *ReporterGitlab) GetProject(client types.HTTPClient, project string, token string) types.RasicProject {
	url := baseURL + apiPath + "projects/" + project

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return types.RasicProject{}
	}

	if res.Status == OK {
		var project types.GitlabProject
		if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
			pterm.Info.Println(res.Body)
			return types.RasicProject{}
		}

		var returnValue types.RasicProject
		returnValue.ID = project.ID
		returnValue.WebURL = project.WebURL
		returnValue.DefaultBranch = project.DefaultBranch

		return returnValue
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}
	return types.RasicProject{}
}

func (a *ReporterGitlab) GetIssues(client types.HTTPClient, project string, token string) []types.RasicIssue {
	url := baseURL + apiPath + "projects/" + project + "/issues?per_page=100"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil
	}

	if res.Status == OK {
		var issuelist types.GitlabIssues
		if err := json.NewDecoder(res.Body).Decode(&issuelist); err != nil {
			return nil
		}

		var returnValue []types.RasicIssue

		for _, issue := range issuelist {
			ele := types.RasicIssue{
				ID:          issue.ID,
				Title:       issue.Title,
				Description: issue.Description,
				State:       issue.State,
			}
			returnValue = append(returnValue, ele)
		}

		return returnValue
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}
	return nil
}

func (a *ReporterGitlab) GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string {
	url := baseURL + apiPath + "projects/" + project + "/repository/files/" + filepath + "/raw?ref=" + fileref

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return ""
	}

	if res.Status == OK {
		fileContent, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			return ""
		}
		return string(fileContent)
	}

	return ""
}

func (a *ReporterGitlab) CreateIssue(client types.HTTPClient, project string, token string, issue types.RasicIssue) types.RasicIssue {
	url := baseURL + apiPath + "projects/" + project + "/issues"

	newGitlabIssue := types.GitlabIssue{
		Title:       issue.Title,
		Description: issue.Description,
		Labels:      issue.Labels,
		CreatedAt:   time.Now(),
	}
	body, marshalErr := json.Marshal(newGitlabIssue)
	if marshalErr != nil {
		pterm.Error.Println(marshalErr)
		return types.RasicIssue{}
	}

	res, err := apiCallPost(client, url, token, string(body))
	if err != nil {
		pterm.Error.Println(err)
		return types.RasicIssue{}
	}

	if res.Status == "201 Created" {
		var issue types.GitlabIssue
		if err := json.NewDecoder(res.Body).Decode(&issue); err != nil {
			return types.RasicIssue{}
		}
		var returnValue types.RasicIssue
		returnValue.ID = issue.ID
		returnValue.Title = issue.Title
		returnValue.Description = issue.Description
		returnValue.State = issue.State

		return returnValue
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}
	return types.RasicIssue{}
}

func (a *ReporterGitlab) GetLabels(client types.HTTPClient, project string, token string) []types.RasicLabel {
	url := baseURL + apiPath + "projects/" + project + "/labels"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil
	}

	if res.Status == OK {
		var labelList []types.GitlabLabel
		if err := json.NewDecoder(res.Body).Decode(&labelList); err != nil {
			return nil
		}

		var returnValue []types.RasicLabel

		for _, label := range labelList {
			ele := types.RasicLabel{
				Name:        label.Name,
				Description: label.Description,
				Color:       label.Color,
				Priority:    0,
			}
			returnValue = append(returnValue, ele)
		}

		return returnValue
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}
	return nil
}

func (a *ReporterGitlab) CreateLabel(client types.HTTPClient, project string, token string, label types.RasicLabel) types.RasicLabel {
	url := baseURL + apiPath + "projects/" + project + "/labels"

	newGitlabLabel := types.GitlabLabel{
		Name:  label.Name,
		Color: label.Color,
	}
	body, marshalErr := json.Marshal(newGitlabLabel)
	if marshalErr != nil {
		pterm.Error.Println(marshalErr)
		return types.RasicLabel{}
	}

	res, err := apiCallPost(client, url, token, string(body))
	if err != nil {
		pterm.Error.Println(err)
		return types.RasicLabel{}
	}

	if res.Status == "201 Created" {
		var label types.GitlabLabel
		if err := json.NewDecoder(res.Body).Decode(&label); err != nil {
			return types.RasicLabel{}
		}
		var returnValue types.RasicLabel
		returnValue.Name = label.Name
		returnValue.Color = label.Color

		return returnValue
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}
	return types.RasicLabel{}
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "REPORTER_PLUGIN",
	MagicCookieValue: "allow",
}

// register types to gob
// this is required to proper serialize and deserialize the data
func init() {
	gob.Register(http.DefaultClient)
	gob.Register(types.RasicIssue{})
	gob.Register(types.RasicLabel{})
	gob.Register(map[string]interface{}{})
}

func main() {
	gitlab := &ReporterGitlab{}

	var pluginMap = map[string]plugin.Plugin{
		"gitlab": &plugins.ReporterPlugin{Impl: gitlab},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

// do a get api call against gitlab.com
func apiCallGet(client types.HTTPClient, url string, token string) (*http.Response, error) {
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
func apiCallPost(client types.HTTPClient, url string, token string, body string) (*http.Response, error) {
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
