package plugins

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"gitlab.com/jstang/rasic/types"
)

type APIRPC struct{ client *rpc.Client }

type APIRPCServer struct {
	Impl API
}

type APIPlugin struct {
	Impl API
}

func (p *APIPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &APIRPCServer{Impl: p.Impl}, nil
}

func (APIPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &APIRPC{client: c}, nil
}

type API interface {
	GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject
	GetProject(client types.HTTPClient, project string, token string) types.RasicProject
	GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string
	GetRepositories(client types.HTTPClient, project string, token string) []types.RasicRepository
	GetRepository(client types.HTTPClient, repository string, token string) types.RasicRepository
}

// GetProjects
func (g *APIRPC) GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject {
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

func (s *APIRPCServer) GetProjects(args map[string]interface{}, resp *[]types.RasicProject) error {
	*resp = s.Impl.GetProjects(args["client"].(types.HTTPClient), args["group"].(string), args["token"].(string))
	return nil
}

// GetProject
func (g *APIRPC) GetProject(client types.HTTPClient, project string, token string) types.RasicProject {
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

func (s *APIRPCServer) GetProject(args map[string]interface{}, resp *types.RasicProject) error {
	*resp = s.Impl.GetProject(args["client"].(types.HTTPClient), args["project"].(string), args["token"].(string))
	return nil
}

// GetFile
func (g *APIRPC) GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string {
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

func (s *APIRPCServer) GetFile(args map[string]interface{}, resp *string) error {
	*resp = s.Impl.GetFile(args["client"].(types.HTTPClient), args["project"].(string), args["filepath"].(string), args["fileref"].(string), args["token"].(string))
	return nil
}

// GetRepositories
func (g *APIRPC) GetRepositories(client types.HTTPClient, project string, token string) []types.RasicRepository {
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

func (s *APIRPCServer) GetRepositories(args map[string]interface{}, resp *[]types.RasicRepository) error {
	*resp = s.Impl.GetRepositories(args["client"].(types.HTTPClient), args["project"].(string), args["token"].(string))
	return nil
}

// GetRepository
func (g *APIRPC) GetRepository(client types.HTTPClient, repository string, token string) types.RasicRepository {
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

func (s *APIRPCServer) GetRepository(args map[string]interface{}, resp *types.RasicRepository) error {
	*resp = s.Impl.GetRepository(args["client"].(types.HTTPClient), args["repository"].(string), args["token"].(string))
	return nil
}
