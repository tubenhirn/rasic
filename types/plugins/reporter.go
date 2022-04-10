package plugins

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"gitlab.com/jstang/rasic/types"
)

type ReporterRPC struct{ client *rpc.Client }

type ReporterRPCServer struct {
	Impl Reporter
}

type ReporterPlugin struct {
	Impl Reporter
}

func (p *ReporterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ReporterRPCServer{Impl: p.Impl}, nil
}

func (ReporterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ReporterRPC{client: c}, nil
}

type Reporter interface {
	GetProjects(client types.HttpClient, group string, token string) []types.RasicProject
	GetProject(client types.HttpClient, project string, token string) types.RasicProject
	GetIssues(client types.HttpClient, project string, token string) []types.RasicIssue
	GetFile(client types.HttpClient, project string, filepath string, fileref string, token string) string
	CreateIssue(client types.HttpClient, project string, token string, issue types.RasicIssue) types.RasicIssue
}

// GetProjects
func (g *ReporterRPC) GetProjects(client types.HttpClient, group string, token string) []types.RasicProject {
	var resp []types.RasicProject
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

func (s *ReporterRPCServer) GetProjects(args map[string]interface{}, resp *[]types.RasicProject) error {
	*resp = s.Impl.GetProjects(args["client"].(types.HttpClient), args["group"].(string), args["token"].(string))
	return nil
}

//GetProject
func (g *ReporterRPC) GetProject(client types.HttpClient, project string, token string) types.RasicProject {
	var resp types.RasicProject
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

func (s *ReporterRPCServer) GetProject(args map[string]interface{}, resp *types.RasicProject) error {
	*resp = s.Impl.GetProject(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string))
	return nil
}

//GetIssues
func (g *ReporterRPC) GetIssues(client types.HttpClient, project string, token string) []types.RasicIssue {
	var resp []types.RasicIssue
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

func (s *ReporterRPCServer) GetIssues(args map[string]interface{}, resp *[]types.RasicIssue) error {
	*resp = s.Impl.GetIssues(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string))
	return nil
}

//GetFile
func (g *ReporterRPC) GetFile(client types.HttpClient, project string, filepath string, fileref string, token string) string {
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

func (s *ReporterRPCServer) GetFile(args map[string]interface{}, resp *string) error {
	*resp = s.Impl.GetFile(args["client"].(types.HttpClient), args["project"].(string), args["filepath"].(string), args["fileref"].(string), args["token"].(string))
	return nil
}

//CreateIssue
func (g *ReporterRPC) CreateIssue(client types.HttpClient, project string, token string, issue types.RasicIssue) types.RasicIssue {
	var resp types.RasicIssue
	err := g.client.Call("Plugin.CreateIssue", map[string]interface{}{
		"client":  client,
		"project": project,
		"token":   token,
		"issue":   issue,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ReporterRPCServer) CreateIssue(args map[string]interface{}, resp *types.RasicIssue) error {
	*resp = s.Impl.CreateIssue(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string), args["issue"].(types.RasicIssue))
	return nil
}
