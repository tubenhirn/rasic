package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"io"

	"github.com/hashicorp/go-plugin"
	"github.com/pterm/pterm"
	"github.com/urfave/cli/v2"
	"gitlab.com/jstang/rasic/types"
	"gitlab.com/jstang/rasic/types/plugins"
)

const baseURL = "https://gitlab.com"
const apiPath = "/api/v4/"
const projectPath = "projects/"

const OK = "200 OK"

type ReporterGitlab struct{}

func (a *ReporterGitlab) GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject {
	// url := baseURL + apiPath + "groups/" + group + "/projects?per_page=100&include_subgroups=true&archived=false"
	url := fmt.Sprintf("%s%sgroups/%s/projects?per_page=100&include_subgroups=true&archived=false", baseURL, apiPath, group)

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
	_, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		pterm.Error.Println(errRead)
	}

	defer res.Body.Close()

	return nil
}

func (a *ReporterGitlab) GetProject(client types.HTTPClient, project string, token string) types.RasicProject {
	// url := baseURL + apiPath + projectPath + project
	url := fmt.Sprintf("%s%s%s%s", baseURL, apiPath, projectPath, project)

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
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return types.RasicProject{}
}

func pagination(client types.HTTPClient, url string, token string, collectedIssues *[]types.RasicIssue, page int) []types.RasicIssue {
	url = fmt.Sprintf("%s&page=%d", url, page)

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

		for _, issue := range issuelist {
			ele := types.RasicIssue{
				ID:          issue.ID,
				Title:       issue.Title,
				Description: issue.Description,
				State:       issue.State,
			}
			*collectedIssues = append(*collectedIssues, ele)
		}

		// look for next page header
		// if set call function again
		// X-Next-Page is a gitlab.com specific pagination header
		if res.Header.Get("X-Next-Page") != "" {
			nextPage := page + 1
			pagination(client, url, token, collectedIssues, nextPage)
		}

		return *collectedIssues
	}

	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return nil
}

// get issues of a group or project
func (a *ReporterGitlab) GetIssues(client types.HTTPClient, subject string, subjectID string, token string) []types.RasicIssue {
	pageSize := 20
	startPage := 1

	// filter issues by attributes
	// TODO: make those filters a config
	state := "opened"
	labels := "cve"

	// url := baseURL + apiPath + subject + "/" + subjectID + "/issues?per_page=" + strconv.Itoa(pageSize) + "&labels=" + labels + "&state=" + state
	url := fmt.Sprintf("%s%s%s/%s/issues?per_page=%d&labels=%s&state=%s", baseURL, apiPath, subject, subjectID, pageSize, labels, state)

	var collectedIssues []types.RasicIssue
	collectedIssues = pagination(client, url, token, &collectedIssues, startPage)

	return collectedIssues
}

// edit a issue of a project
func (a *ReporterGitlab) EditIssue(client types.HTTPClient, projectID string, issueID string, token string, editPayload types.RasicIssueUpdate) types.RasicIssue {
	// url := baseURL + apiPath + projectPath + projectID + "/issues/" + issueID
	url := fmt.Sprintf("%s%s%s/%s/issues/%s", baseURL, apiPath, projectPath, projectID, issueID)

	issueUpdate := types.GitlabIssueUpdate{
		StateEvent: editPayload.State,
	}

	res, err := apiCallPut(client, url, token, issueUpdate)
	if err != nil {
		pterm.Error.Println(err)
		return types.RasicIssue{}
	}

	if res.Status == OK {
		var issue types.GitlabIssue
		if err := json.NewDecoder(res.Body).Decode(&issue); err != nil {
			return types.RasicIssue{}
		}
		var returnValue types.RasicIssue
		returnValue.ID = issue.ID
		returnValue.Title = issue.Title
		returnValue.State = issue.State

		return returnValue
	}

	defer res.Body.Close()

	return types.RasicIssue{}
}

func (a *ReporterGitlab) GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string {
	// url := baseURL + apiPath + projectPath + project + "/repository/files/" + filepath + "/raw?ref=" + fileref
	url := fmt.Sprintf("%s%s%s/%s/repository/files/%s/raw?ref=%s", baseURL, apiPath, projectPath, project, filepath, fileref)

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

func (a *ReporterGitlab) CreateIssue(client types.HTTPClient, project string, token string, issue types.RasicIssue) types.RasicIssue {
	// url := baseURL + apiPath + projectPath + project + "/issues"
	url := fmt.Sprintf("%s%s%s/%s/issues", baseURL, apiPath, projectPath, project)

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
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return types.RasicIssue{}
}

func (a *ReporterGitlab) GetLabels(client types.HTTPClient, project string, token string) []types.RasicLabel {
	// url := baseURL + apiPath + projectPath + project + "/labels"
	url := fmt.Sprintf("%s%s%s/%s/labels", baseURL, apiPath, projectPath, project)

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
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

	return nil
}

func (a *ReporterGitlab) CreateLabel(client types.HTTPClient, project string, token string, label types.RasicLabel) types.RasicLabel {
	// url := baseURL + apiPath + projectPath + project + "/labels"
	url := fmt.Sprintf("%s%s%s/%s/labels", baseURL, apiPath, projectPath, project)

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
	_, err = io.ReadAll(res.Body)
	if err != nil {
		pterm.Error.Println(err)
	}

	defer res.Body.Close()

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
	gitlab := &ReporterGitlab{}

	var pluginMap = map[string]plugin.Plugin{
		"reporter_gitlab": &plugins.ReporterPlugin{Impl: gitlab},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}

// do a get api call against gitlab.com
func apiCallGet(client types.HTTPClient, url string, token string) (*http.Response, error) {
	req, reqerr := http.NewRequest(http.MethodGet, url, nil)

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
	req, reqerr := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

	if reqerr != nil {
		return nil, cli.Exit(reqerr, 1)
	}

	// set auth header
	req.Header.Set("PRIVATE-TOKEN", token)
	// set content type
	req.Header.Set("content-type", "application/json")

	// do the request
	res, reserr := client.Do(req)
	if reserr != nil {
		return nil, cli.Exit(reserr, 1)
	}

	return res, nil
}

// do a put api call against gitlab.com
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
	req.Header.Set("PRIVATE-TOKEN", token)
	// set content type
	req.Header.Set("content-type", "application/json")

	// do the request
	res, reserr := client.Do(req)
	if reserr != nil {
		return nil, cli.Exit(reserr, 1)
	}

	return res, nil
}
