package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/safedep/dry/adapters"
)

type githubToolProvider struct {
	client *adapters.GithubClient
}

// ResolveRefToSha resolves a reference to a commit SHA.
func (p *githubToolProvider) ResolveRefToSha(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner, ok := request.Params.Arguments["owner"].(string)
	if !ok {
		return nil, fmt.Errorf("owner must be a string")
	}

	repo, ok := request.Params.Arguments["repo"].(string)
	if !ok {
		return nil, fmt.Errorf("repo must be a string")
	}

	ref, ok := request.Params.Arguments["ref"].(string)
	if !ok {
		return nil, fmt.Errorf("ref must be a string")
	}

	commit, _, err := p.client.Client.Repositories.GetCommit(ctx, owner, repo, ref, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get commit: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%s %s", ref, commit.GetSHA())), nil
}

// GetLatestPinnedVersion returns the latest pinned version for a given repository.
// Latest version is determined by the latest non-prerelease version.
func (p *githubToolProvider) GetLatestPinnedVersion(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner, ok := request.Params.Arguments["owner"].(string)
	if !ok {
		return nil, fmt.Errorf("owner must be a string")
	}

	repo, ok := request.Params.Arguments["repo"].(string)
	if !ok {
		return nil, fmt.Errorf("repo must be a string")
	}

	latestRelease, _, err := p.client.Client.Repositories.GetLatestRelease(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest release: %w", err)
	}

	commit, _, err := p.client.Client.Repositories.GetCommit(ctx, owner, repo, latestRelease.GetTagName(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get commit: %w", err)
	}

	return mcp.NewToolResultText(fmt.Sprintf("%s %s", latestRelease.GetTagName(), commit.GetSHA())), nil
}

// GetVersions returns the latest versions (releases) for a given repository.
func (p *githubToolProvider) GetVersions(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	owner, ok := request.Params.Arguments["owner"].(string)
	if !ok {
		return nil, fmt.Errorf("owner must be a string")
	}

	repo, ok := request.Params.Arguments["repo"].(string)
	if !ok {
		return nil, fmt.Errorf("repo must be a string")
	}

	releases, _, err := p.client.Client.Repositories.ListReleases(ctx, owner, repo, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list releases: %w", err)
	}

	var versions []string
	for _, release := range releases {
		if release.GetTagName() != "" {
			versions = append(versions, release.GetTagName())
		}
		if len(versions) >= 10 { // Limit to 10
			break
		}
	}

	return mcp.NewToolResultText(strings.Join(versions, "\n")), nil
}

func RegisterGithubTool(srv *server.MCPServer) error {
	client, err := adapters.NewGithubClient(adapters.DefaultGitHubClientConfig())
	if err != nil {
		return fmt.Errorf("failed to create github client: %w", err)
	}

	resolveRefToShaTool := mcp.NewTool("github_resolve_ref_to_sha",
		mcp.WithDescription("Resolve a Github reference such as a branch or tag to a commit SHA. Returns the ref and SHA for pinning GitHub Actions."),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("The owner of the repository"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("The name of the repository"),
		),
		mcp.WithString("ref",
			mcp.Required(),
			mcp.Description("The reference to resolve"),
		),
	)

	getLatestPinnedVersionTool := mcp.NewTool("github_get_latest_pinned_version",
		mcp.WithDescription("Get the latest pinned version and its tag for a given repository. Used to update pinned GitHub Actions"),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("The owner of the repository"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("The name of the repository"),
		),
	)

	getVersionsTool := mcp.NewTool("github_get_versions",
		mcp.WithDescription("List the latest versions (releases) for a given repository. Used to check available updates for GitHub Actions"),
		mcp.WithString("owner",
			mcp.Required(),
			mcp.Description("The owner of the repository"),
		),
		mcp.WithString("repo",
			mcp.Required(),
			mcp.Description("The name of the repository"),
		),
	)

	toolProvider := &githubToolProvider{client: client}

	srv.AddTool(resolveRefToShaTool, toolProvider.ResolveRefToSha)
	srv.AddTool(getLatestPinnedVersionTool, toolProvider.GetLatestPinnedVersion)
	srv.AddTool(getVersionsTool, toolProvider.GetVersions)

	return nil
}
