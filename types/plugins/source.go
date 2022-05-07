package plugins

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"gitlab.com/jstang/rasic/types"
)

type SourceRPC struct{ client *rpc.Client }

type SourceRPCServer struct {
	Impl Source
}

type SourcePlugin struct {
	Impl Source
}

func (p *SourcePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &SourceRPCServer{Impl: p.Impl}, nil
}

func (SourcePlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &SourceRPC{client: c}, nil
}

type Source interface {
	GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject
	GetProject(client types.HTTPClient, project string, token string) types.RasicProject
	GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string
	GetRepositories(client types.HTTPClient, project string, token string) []types.RasicRepository
	GetRepository(client types.HTTPClient, repository string, token string) types.RasicRepository
}

// GetProjects
func (g *SourceRPC) GetProjects(client types.HTTPClient, group string, token string) []types.RasicProject {
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

func (s *SourceRPCServer) GetProjects(args map[string]interface{}, resp *[]types.RasicProject) error {
	*resp = s.Impl.GetProjects(args["client"].(types.HTTPClient), args["group"].(string), args["token"].(string))
	return nil
}

// GetProject
func (g *SourceRPC) GetProject(client types.HTTPClient, project string, token string) types.RasicProject {
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

func (s *SourceRPCServer) GetProject(args map[string]interface{}, resp *types.RasicProject) error {
	*resp = s.Impl.GetProject(args["client"].(types.HTTPClient), args["project"].(string), args["token"].(string))
	return nil
}

// GetFile
func (g *SourceRPC) GetFile(client types.HTTPClient, project string, filepath string, fileref string, token string) string {
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

func (s *SourceRPCServer) GetFile(args map[string]interface{}, resp *string) error {
	*resp = s.Impl.GetFile(args["client"].(types.HTTPClient), args["project"].(string), args["filepath"].(string), args["fileref"].(string), args["token"].(string))
	return nil
}

// GetRepositories
func (g *SourceRPC) GetRepositories(client types.HTTPClient, project string, token string) []types.RasicRepository {
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

func (s *SourceRPCServer) GetRepositories(args map[string]interface{}, resp *[]types.RasicRepository) error {
	*resp = s.Impl.GetRepositories(args["client"].(types.HTTPClient), args["project"].(string), args["token"].(string))
	return nil
}

// GetRepository
func (g *SourceRPC) GetRepository(client types.HTTPClient, repository string, token string) types.RasicRepository {
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

func (s *SourceRPCServer) GetRepository(args map[string]interface{}, resp *types.RasicRepository) error {
	*resp = s.Impl.GetRepository(args["client"].(types.HTTPClient), args["repository"].(string), args["token"].(string))
	return nil
}
