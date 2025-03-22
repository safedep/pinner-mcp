package tools

import (
	"context"
	"fmt"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type dockerToolProvider struct{}

// ResolveImageToDigest resolves an image to a digest.
func (p *dockerToolProvider) ResolveImageToDigest(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	image, ok := request.Params.Arguments["image"].(string)
	if !ok {
		return nil, fmt.Errorf("image is required")
	}

	version, ok := request.Params.Arguments["version"].(string)
	if !ok {
		return nil, fmt.Errorf("version is required")
	}

	ref, err := parseContainerImageName(image, version)
	if err != nil {
		return nil, fmt.Errorf("failed to parse image name: %w", err)
	}

	desc, err := remote.Get(ref, remote.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch image digest: %w", err)
	}

	return mcp.NewToolResultText(desc.Digest.String()), nil
}

func RegisterDockerTool(srv *server.MCPServer) error {
	resolveImageToDigestTool := mcp.NewTool(
		"docker_resolve_image_to_digest",
		mcp.WithDescription("Resolve a container image version to a digest for pinning to immutable images. Use to resolve base images in Dockerfile."),
		mcp.WithString("image",
			mcp.Required(),
			mcp.Description("The image to resolve"),
		),
		mcp.WithString("version",
			mcp.Required(),
			mcp.Description("The version of the image to resolve"),
		),
	)

	toolProvider := &dockerToolProvider{}
	srv.AddTool(resolveImageToDigestTool, toolProvider.ResolveImageToDigest)

	return nil
}

func parseContainerImageName(image, version string) (name.Reference, error) {
	fullName := fmt.Sprintf("%s:%s", image, version)

	ref, err := name.ParseReference(fullName)
	if err != nil {
		return nil, fmt.Errorf("failed to parse image name: %w", err)
	}

	return ref, nil
}
