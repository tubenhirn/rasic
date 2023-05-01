package main

import (
	"encoding/gob"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

const baseURL = "https://gitlab.com"
const apiPath = "/api/v4/"

const OK = "200 OK"

type SourceGitlab struct{}

func (a *SourceGitlab) GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject {
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
				FullName: "",
				ProjectType: "gitlab",
			}
			returnValue = append(returnValue, ele)
		}

		return returnValue
	}
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return nil
}

func (a *SourceGitlab) GetProject(client types.HTTPClient, project string, token string) types.RasicProject {
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
		returnValue.FullName = ""
		returnValue.ProjectType = "gitlab"

		return returnValue
	}
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return types.RasicProject{}
}

func (a *SourceGitlab) GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string {
	url := baseURL + apiPath + "projects/" + project + "/repository/files/" + filepath + "/raw?ref=" + fileref

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return ""
	}

	if res.Status == OK {
		fileContent, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return ""
		}
		return string(fileContent)
	}

	defer res.Body.Close()

	return ""
}

func (a *SourceGitlab) GetRepositories(client types.HTTPClient, project string, token string) []types.RasicRepository {
	url := baseURL + apiPath + "projects/" + project + "/registry/repositories"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return nil
	}

	if res.Status == OK {
		var repositorylist types.GitlabRepositories
		if err := json.NewDecoder(res.Body).Decode(&repositorylist); err != nil {
			return nil
		}

		var returnValue []types.RasicRepository

		for _, repo := range repositorylist {
			ele := types.RasicRepository{
				ID: repo.ID,
			}
			returnValue = append(returnValue, ele)
		}
		return returnValue
	}
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return nil
}

func (a *SourceGitlab) GetRepository(client types.HTTPClient, repository string, token string) types.RasicRepository {
	url := baseURL + apiPath + "registry/repositories/" + repository + "?tags=true"

	res, err := apiCallGet(client, url, token)

	if err != nil {
		pterm.Error.Println(err)
		return types.RasicRepository{}
	}

	if res.Status == OK {
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
			ID:  repo.ID,
			Tag: latestTag,
		}

		return returnValue
	}
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return types.RasicRepository{}
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SOURCE_PLUGIN",
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
	gitlab := &SourceGitlab{}

	var pluginMap = map[string]plugin.Plugin{
		"source_gitlab": &plugins.SourcePlugin{Impl: gitlab},
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
		return nil, cli.Exit(reqerr, 1)
	}

	// set auth header
	req.Header.Set("PRIVATE-TOKEN", token)

	// do the request
	res, reserr := client.Do(req)
	if reserr != nil {
		return nil, cli.Exit(reserr, 1)
	}

	// retrun the response
	return res, nil
}

// do a post api call against gitlab.com
func apiCallPost(client types.HTTPClient, url string, token string, body string) (*http.Response, error) {
	req, reqerr := http.NewRequest("POST", url, strings.NewReader(body))

	if reqerr != nil {
		return nil, cli.Exit(reqerr, 1)
	}

	// set auth header
	req.Header.Set("PRIVATE-TOKEN", token)
	req.Header.Set("content-type", "application/json")

	// do the request
	res, reserr := client.Do(req)
	if reserr != nil {
		return nil, cli.Exit(reserr, 1)
	}

	return res, nil
}
