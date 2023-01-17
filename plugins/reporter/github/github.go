package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

const baseURI = "https://api.github.com"
const apiVersion = "2022-11-28"

const OK = "200 OK"

type ReporterGithub struct{}

func (a *ReporterGithub) GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject {

	return nil
}

func (a *ReporterGithub) GetProject(client types.HTTPClient, project string, token string) types.RasicProject {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, "repos", project)

	res, err := apiCallGet(client, uri.String(), token)

	if err != nil {
		return types.RasicProject{}
	}

	if res.Status == OK {
		var project types.GithubRepositorie
		if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
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

// get issues of a group or project
func (a *ReporterGithub) GetIssues(client types.HTTPClient, subject string, subjectID string, token string) []types.RasicIssue {

	return nil
}

// edit a issue of a project
func (a *ReporterGithub) EditIssue(client types.HTTPClient, projectID string, issueID string, token string, editPayload types.RasicIssueUpdate) types.RasicIssue {

	return types.RasicIssue{}
}

func (a *ReporterGithub) GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string {
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

func (a *ReporterGithub) CreateIssue(client types.HTTPClient, project string, token string, issue types.RasicIssue) types.RasicIssue {

	return types.RasicIssue{}
}

func (a *ReporterGithub) GetLabels(client types.HTTPClient, project string, token string) []types.RasicLabel {

	return nil
}

func (a *ReporterGithub) CreateLabel(client types.HTTPClient, project string, token string, label types.RasicLabel) types.RasicLabel {

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
	gob.Register(types.RasicIssueUpdate{})
	gob.Register(types.RasicLabel{})
	gob.Register(map[string]interface{}{})
}

func main() {
	gitlab := &ReporterGithub{}

	var pluginMap = map[string]plugin.Plugin{
		"reporter_github": &plugins.ReporterPlugin{Impl: gitlab},
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

// do a post api call against github.com
func apiCallPost(client types.HTTPClient, url string, token string, body string) (*http.Response, error) {
	req, reqerr := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

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

	return res, nil
}

// do a put api call against github.com
func apiCallPut(client types.HTTPClient, url string, token string, payload interface{}) (*http.Response, error) {
	reqPayload, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return nil, cli.Exit(marshalErr, 1)
	}

	req, reqerr := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(reqPayload))
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

	return res, nil
}
