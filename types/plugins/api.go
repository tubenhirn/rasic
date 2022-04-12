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
	GetProjects(client types.HttpClient, group string, token string) []types.RasicProject
	GetProject(client types.HttpClient, project string, token string) types.RasicProject
	GetFile(client types.HttpClient, project string, filepath string, fileref string, token string) string
	GetRepositories(client types.HttpClient, project string, token string) []types.RasicRepository
	GetRepository(client types.HttpClient, repository string, token string) types.RasicRepository
}

// GetProjects
func (g *ApiRPC) GetProjects(client types.HttpClient, group string, token string) []types.RasicProject {
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

func (s *ApiRPCServer) GetProjects(args map[string]interface{}, resp *[]types.RasicProject) error {
	*resp = s.Impl.GetProjects(args["client"].(types.HttpClient), args["group"].(string), args["token"].(string))
	return nil
}

//GetProject
func (g *ApiRPC) GetProject(client types.HttpClient, project string, token string) types.RasicProject {
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

func (s *ApiRPCServer) GetProject(args map[string]interface{}, resp *types.RasicProject) error {
	*resp = s.Impl.GetProject(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string))
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

//GetRepositories
func (g *ApiRPC) GetRepositories(client types.HttpClient, project string, token string) []types.RasicRepository {
	var resp []types.RasicRepository
	err := g.client.Call("Plugin.GetRepositories", map[string]interface{}{
		"client":  client,
		"project": project,
		"token":   token,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ApiRPCServer) GetRepositories(args map[string]interface{}, resp *[]types.RasicRepository) error {
	*resp = s.Impl.GetRepositories(args["client"].(types.HttpClient), args["project"].(string), args["token"].(string))
	return nil
}

//GetRepository
func (g *ApiRPC) GetRepository(client types.HttpClient, repository string, token string) types.RasicRepository {
	var resp types.RasicRepository
	err := g.client.Call("Plugin.GetRepository", map[string]interface{}{
		"client":     client,
		"repository": repository,
		"token":      token,
	}, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (s *ApiRPCServer) GetRepository(args map[string]interface{}, resp *types.RasicRepository) error {
	*resp = s.Impl.GetRepository(args["client"].(types.HttpClient), args["repository"].(string), args["token"].(string))
	return nil
}
