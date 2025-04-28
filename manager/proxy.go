package manager

import (
	"fmt"
	"strings"

	"github.com/mymcp/mcp-proxy/pkg/common"
	"github.com/mymcp/mcp-proxy/pkg/proxy"
	"github.com/tk103331/mymcp/manager/data"
)

func newProxyServer(cfg *data.ServerConfig) (*proxy.ProxyServer, error) {
	commandLine := cfg.CommandLine
	environment := cfg.Environment
	envs := []string{}
	for paramName, paramValue := range cfg.Parameters {
		commandLine = strings.ReplaceAll(commandLine, "${"+paramName+"}", paramValue)
		for envName, envValue := range environment {
			envs = append(envs, fmt.Sprintf("%s=%s", envName, strings.ReplaceAll(envValue, "${"+paramName+"}", paramValue)))
		}
	}

	fields := strings.Fields(commandLine)
	mcpConfig := common.McpServerConfig{
		Transport: "stdio",
		Cmd:       fields[0],
		Args:      fields[1:],
		Env:       envs,
	}
	server, err := proxy.NewProxyServer(&mcpConfig)
	if err != nil {
		return nil, err
	}
	return server, nil
}
