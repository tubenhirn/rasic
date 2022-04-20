package types

import (
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-plugin"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type RasicIssue struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string
	Severity    Severity
}

type RasicProject struct {
	Id             int
	WebUrl         string
	DefaultBranch  string
	IgnoreFileName string
}

type RasicRepository struct {
	Id  int
	Tag RasicTag
}

type RasicTag struct {
	Location string
	Name     string
	Path     string
}

type RasicPlugin struct {
	PluginHome   string
	PluginPath   string
	PluginName   string
	PluginConfig plugin.HandshakeConfig
	PluginMap    map[string]plugin.Plugin
}

type Severity int64

const (
	Unknown Severity = iota
	Low
	Medium
	High
	Critical
)

func (s Severity) String() string {
	return [...]string{"UNKNOWN", "LOW", "MEDIUM", "HIGH", "CRITICAL"}[s]
}

func (s *Severity) FromString(severity string) Severity {
	return map[string]Severity{
		"UNKNOWN":  Unknown,
		"LOW":      Low,
		"MEDIUM":   Medium,
		"HIGH":     High,
		"CRITICAL": Critical,
	}[severity]
}

func (s Severity) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *Severity) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	*s = s.FromString(str)
	return nil
}
