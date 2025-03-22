package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.com/safedep/dry/log"
	"github.com/safedep/pinner-mcp/tools"
)

func registerTools(srv *server.MCPServer) error {
	if err := tools.RegisterGithubTool(srv); err != nil {
		return fmt.Errorf("failed to register github tool: %w", err)
	}

	return nil
}

func main() {
	log.Init("pinner-mcp", "dev")

	srv := server.NewMCPServer("pinner-mcp", "0.0.1",
		server.WithInstructions("This is a Model Context Protocol (MCP) server that can help pin GitHub Actions to a specific commit hash."),
		server.WithLogging(),
	)

	if err := registerTools(srv); err != nil {
		log.Fatalf("failed to register tools: %v", err)
	}

	if err := server.ServeStdio(srv); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
