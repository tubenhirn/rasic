package main

import (
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

const baseURI = "https://api.github.com"
const apiVersion = "2022-11-28"

const OK = "200 OK"

type SourceGithub struct{}

// register types to gob
// this is required to proper serialize and deserialize the data
func init() {
	gob.Register(http.DefaultClient)
	gob.Register(types.RasicIssue{})
	gob.Register(types.RasicRepository{})
	gob.Register(map[string]interface{}{})
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SOURCE_PLUGIN",
	MagicCookieValue: "allow",
}

func (a *SourceGithub) GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject {

	return nil
}

func (a *SourceGithub) GetProject(client types.HTTPClient, project string, token string) types.RasicProject {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, "repos", project)

	res, err := apiCallGet(client, uri.String(), token)

	if err != nil {
		pterm.Error.Println(err)
		return types.RasicProject{}
	}

	if res.Status == OK {
		var project types.GithubRepositorie
		if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
			pterm.Info.Println(res.Body)
			return types.RasicProject{}
		}

		var returnValue types.RasicProject
		returnValue.ID = project.ID
		returnValue.WebURL = project.HtmlUrl
		returnValue.DefaultBranch = project.DefaultBranch
		returnValue.FullName = project.FullName
		returnValue.ProjectType = "github"

		return returnValue
	}
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return types.RasicProject{}
}

func (a *SourceGithub) GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, "repos", project, "contents", filepath)
	res, err := apiCallGet(client, uri.String(), token)

	if err != nil {
		return ""
	}

	if res.Status == OK {
		var file types.GithubFile
		if err := json.NewDecoder(res.Body).Decode(&file); err != nil {
			return ""
		}
		fileContent, err := base64.StdEncoding.DecodeString(file.Content)
		if err != nil {
			return ""
		}
		return string(fileContent)
	}

	defer res.Body.Close()

	return ""
}

func (a *SourceGithub) GetRepositories(client types.HTTPClient, project string, token string) []types.RasicRepository {

	return nil
}

func (a *SourceGithub) GetRepository(client types.HTTPClient, repository string, token string) types.RasicRepository {

	return types.RasicRepository{}
}

func main() {
	github := &SourceGithub{}

	var pluginMap = map[string]plugin.Plugin{
		"source_github": &plugins.SourcePlugin{Impl: github},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

// do a get api call against github.com
func apiCallGet(client types.HTTPClient, url string, token string) (*http.Response, error) {
	req, reqerr := http.NewRequest("GET", url, nil)

	if reqerr != nil {
		return nil, cli.Exit(reqerr, 1)
	}

	// set auth header
	basic := base64.StdEncoding.EncodeToString([]byte(token))
	req.Header.Add("Authorization", "Basic "+basic)

	// set github api version
	req.Header.Add("X-GitHub-Api-Version", apiVersion)

	// set content header
	req.Header.Add("Accept", "application/vnd.github+json")

	// do the request
	res, reserr := client.Do(req)
	if reserr != nil {
		return nil, cli.Exit(reserr, 1)
	}

	// retrun the response
	return res, nil
}
