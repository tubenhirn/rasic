package plugins

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"gitlab.com/jstang/rasic/types"
)

type ApiRPC struct{ client *rpc.Client }

type ApiRPCServer struct {
	Impl Api
}

type ApiPlugin struct {
	Impl Api
}

func (p *ApiPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ApiRPCServer{Impl: p.Impl}, nil
}

func (ApiPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ApiRPC{client: c}, nil
}

type Api interface {
	GetProjects(client types.HttpClient, group string, token string) types.GitlabProjects
	GetProject(client types.HttpClient, project string, token string) types.GitlabProject
	GetIssues(client types.HttpClient, project string, token string) types.GitlabIssues
	GetFile(client types.HttpClient, project string, filepath string, fileref string, token string) string
	CreateIssue(client types.HttpClient, project string, token string, issue types.Issue) types.GitlabIssue
}

// GetProjects
func (g *ApiRPC) GetProjects(client types.HttpClient, group string, token string) types.GitlabProjects {
	var resp types.GitlabProjects
	err := g.client.Call("Plugin.GetProjects", map[string]interface{}{
		"client": client,
		"group":  group,
		"token":  token,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ApiRPCServer) GetProjects(args map[string]interface{}, resp *types.GitlabProjects) error {
	*resp = s.Impl.GetProjects(args["client"].(types.HttpClient), args["group"].(string), args["token"].(string))
	return nil
}

//GetProject
func (g *ApiRPC) GetProject(client types.HttpClient, project string, token string) types.GitlabProject {
	var resp types.GitlabProject
	err := g.client.Call("Plugin.GetProject", map[string]interface{}{
		"client":  client,
		"project": project,
		"token":   token,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ApiRPCServer) GetProject(args map[string]interface{}, resp *types.GitlabProject) error {
	*resp = s.Impl.GetProject(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string))
	return nil
}

//GetIssues
func (g *ApiRPC) GetIssues(client types.HttpClient, project string, token string) types.GitlabIssues {
	var resp types.GitlabIssues
	err := g.client.Call("Plugin.GetIssues", map[string]interface{}{
		"client":  client,
		"project": project,
		"token":   token,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ApiRPCServer) GetIssues(args map[string]interface{}, resp *types.GitlabIssues) error {
	*resp = s.Impl.GetIssues(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string))
	return nil
}

//GetFile
func (g *ApiRPC) GetFile(client types.HttpClient, project string, filepath string, fileref string, token string) string {
	var resp string
	err := g.client.Call("Plugin.GetFile", map[string]interface{}{
		"client":   client,
		"project":  project,
		"filepath": filepath,
		"fileref":  fileref,
		"token":    token,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ApiRPCServer) GetFile(args map[string]interface{}, resp *string) error {
	*resp = s.Impl.GetFile(args["client"].(types.HttpClient), args["project"].(string), args["filepath"].(string), args["fileref"].(string), args["token"].(string))
	return nil
}

//CreateIssue
func (g *ApiRPC) CreateIssue(client types.HttpClient, project string, token string, issue types.Issue) types.GitlabIssue {
	var resp types.GitlabIssue
	err := g.client.Call("Plugin.CreateIssue", map[string]interface{}{
		"client":  client,
		"project": project,
		"token":   token,
		"issue": issue,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ApiRPCServer) CreateIssue(args map[string]interface{}, resp *types.GitlabIssue) error {
	*resp = s.Impl.CreateIssue(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string), args["issue"].(types.Issue))
	return nil
}
