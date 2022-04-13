package main

import (
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

var baseUrl = "https://gitlab.com"
var apiPath = "/api/v4/"

type ApiGitlab struct{}

func (a *ApiGitlab) GetProjects(client types.HttpClient, group string, token string) []types.RasicProject {
	url := baseUrl + apiPath + "groups/" + group + "/projects?per_page=100&include_subgroups=true&archived=false"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil
	}

	if res.Status == "200 OK" {
		var projectlist types.GitlabProjects
		if err := json.NewDecoder(res.Body).Decode(&projectlist); err != nil {
			pterm.Error.Println(err)
			return nil
		}

		var returnValue []types.RasicProject
		for _, pro := range projectlist {
			ele := types.RasicProject{
				Id:            pro.ID,
				WebUrl:        pro.WebURL,
				DefaultBranch: pro.DefaultBranch,
			}
			returnValue = append(returnValue, ele)

		}

		return returnValue
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return nil
		}
		return nil
	}

}

func (a *ApiGitlab) GetProject(client types.HttpClient, project string, token string) types.RasicProject {
	url := baseUrl + apiPath + "projects/" + project

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return types.RasicProject{}
	}

	if res.Status == "200 OK" {
		var project types.GitlabProject
		if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
			pterm.Info.Println(res.Body)
			return types.RasicProject{}
		}

		var returnValue types.RasicProject
		returnValue.Id = project.ID
		returnValue.WebUrl = project.WebURL
		returnValue.DefaultBranch = project.DefaultBranch

		return returnValue

	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return types.RasicProject{}
		}
		return types.RasicProject{}
	}

}

func (a *ApiGitlab) GetIssues(client types.HttpClient, project string, token string) []types.RasicIssue {
	url := baseUrl + apiPath + "projects/" + project + "/issues?per_page=100"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil
	}

	if res.Status == "200 OK" {
		var issuelist types.GitlabIssues
		if err := json.NewDecoder(res.Body).Decode(&issuelist); err != nil {
			return nil
		}

		var returnValue []types.RasicIssue

		for _, issue := range issuelist {
			ele := types.RasicIssue{
				Id:          issue.ID,
				Title:       issue.Title,
				Description: issue.Description,
				State:       issue.State,
			}
			returnValue = append(returnValue, ele)

		}

		return returnValue
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return nil
		}
		return nil
	}
}

func (a *ApiGitlab) GetFile(client types.HttpClient, project string, filepath string, fileref string, token string) string {
	url := baseUrl + apiPath + "projects/" + project + "/repository/files/" + filepath + "/raw?ref=" + fileref

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return ""
	}

	if res.Status == "200 OK" {
		fileContent, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			return ""
		}
		return string(fileContent)
	}

	return ""
}

func (a *ApiGitlab) CreateIssue(client types.HttpClient, project string, token string, issue types.RasicIssue) types.RasicIssue {
	url := baseUrl + apiPath + "projects/" + project + "/issues"

	body, marshalErr := json.Marshal(issue)
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
		returnValue.Id = issue.ID
		returnValue.Title = issue.Title
		returnValue.Description = issue.Description
		returnValue.State = issue.State

		return returnValue
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return types.RasicIssue{}
		}
		return types.RasicIssue{}
	}
}

func (a *ApiGitlab) GetRepositories(client types.HttpClient, project string, token string) []types.RasicRepository {
	url := baseUrl + apiPath + "projects/" + project + "/registry/repositories"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil
	}

	if res.Status == "200 OK" {
		var repositorylist types.GitlabRepositories
		if err := json.NewDecoder(res.Body).Decode(&repositorylist); err != nil {
			return nil
		}

		var returnValue []types.RasicRepository

		for _, repo := range repositorylist {
			ele := types.RasicRepository{
				Id: repo.ID,
			}
			returnValue = append(returnValue, ele)

		}

		return returnValue
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return nil
		}
		return nil
	}
}

func (a *ApiGitlab) GetRepository(client types.HttpClient, repository string, token string) types.RasicRepository {
	url := baseUrl + apiPath + "registry/repositories/" + repository + "?tags=true"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return types.RasicRepository{}
	}

	if res.Status == "200 OK" {
		var repo types.GitlabRepository
		if err := json.NewDecoder(res.Body).Decode(&repo); err != nil {
			return types.RasicRepository{}
		}

		var latestTag types.RasicTag
		if len(repo.Tags) > 0 {
			latestTag.Location = repo.Tags[len(repo.Tags)-1].Location
			latestTag.Name = repo.Tags[len(repo.Tags)-1].Name
			latestTag.Path = repo.Tags[len(repo.Tags)-1].Path
		}

		returnValue := types.RasicRepository{
			Id:  repo.ID,
			Tag: latestTag,
		}

		return returnValue
	} else {
		_, err := ioutil.ReadAll(res.Body)
		if err != nil {
			pterm.Error.Println(err)
			return types.RasicRepository{}
		}
		return types.RasicRepository{}
	}
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "API_PLUGIN",
	MagicCookieValue: "allow",
}

// register types to gob
// this is required to proper serialize and deserialize the data
func init() {
	gob.Register(http.DefaultClient)
	gob.Register(types.RasicIssue{})
	gob.Register(types.RasicRepository{})
	gob.Register(map[string]interface{}{})
}

func main() {
	gitlab := &ApiGitlab{}

	var pluginMap = map[string]plugin.Plugin{
		"gitlab": &plugins.ApiPlugin{Impl: gitlab},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

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
