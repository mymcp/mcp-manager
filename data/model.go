package data

import (
	"fmt"
	"net/url"

	"github.com/mark3labs/mcp-go/mcp"
)

type Workspace struct {
	ID             string                   `json:"id"`
	Name           string                   `json:"name"`
	Status         string                   `json:"status"`
	Description    string                   `json:"description"`
	Enabled        bool                     `json:"enabled"`
	AutoRun        bool                     `json:"autoRun"`
	ManagedClients map[string]ManagedClient `json:"managedClients"`
}

type ManagedClient struct {
	Config string `json:"config"`
}

type ServerInstance struct {
	ID         string              `json:"id"`
	Config     *ServerConfig       `json:"config"`
	Status     string              `json:"status"`
	Error      string              `json:"error"`
	ServerInfo *mcp.Implementation `json:"serverInfo"`
	Endpoint   string              `json:"endpoint"`
}

type ServerConfig struct {
	ID         string            `json:"id"`
	Workspace  string            `json:"workspace"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Parameters map[string]string `json:"parameters"`
	// for stdio
	CommandLine string            `json:"commandLine"` // cmd with args
	Environment map[string]string `json:"environment"`
}

type Settings struct {
	Language string `json:"language"`
	Theme    string `json:"theme"`
	BaseURL  string `json:"baseUrl"`
}

// ParseBaseURL 从BaseURL中解析出主机和端口
func (s *Settings) ParseBaseURL() (host string, port string, err error) {
	if s.BaseURL == "" {
		return "", "", fmt.Errorf("BaseURL is empty")
	}

	u, err := url.Parse(s.BaseURL)
	if err != nil {
		return "", "", fmt.Errorf("invalid BaseURL: %v", err)
	}

	return u.Hostname(), u.Port(), nil
}
