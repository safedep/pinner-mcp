package tools

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
)

func TestParseContainerImageName(t *testing.T) {
	cases := []struct {
		name     string
		image    string
		version  string
		expected string
		err      error
	}{
		{
			name:     "valid image name",
			image:    "docker.io/library/alpine",
			version:  "latest",
			expected: "docker.io/library/alpine:latest",
		},
		{
			name:     "short name",
			image:    "alpine",
			version:  "latest",
			expected: "alpine:latest",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ref, err := parseContainerImageName(c.image, c.version)
			assert.NoError(t, err)
			assert.Equal(t, c.expected, ref.String())
		})
	}
}

func TestResolveImageToDigest(t *testing.T) {
	cases := []struct {
		name     string
		image    string
		version  string
		expected string
		err      error
	}{
		{
			name:     "valid image name",
			image:    "alpine",
			version:  "3.19.7",
			expected: "sha256:e5d0aea7f7d2954678a9a6269ca2d06e06591881161961ea59e974dff3f12377",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			provider := &dockerToolProvider{}
			request := mcp.CallToolRequest{}
			request.Params.Arguments = map[string]interface{}{
				"image":   c.image,
				"version": c.version,
			}

			result, err := provider.ResolveImageToDigest(context.Background(), request)
			if c.err != nil {
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.NoError(t, err)

				textContent, ok := result.Content[0].(mcp.TextContent)
				assert.True(t, ok)
				assert.Equal(t, c.expected, textContent.Text)
			}
		})
	}
}

func TestGetImageVersions(t *testing.T) {
	cases := []struct {
		name             string
		image            string
		expectedMinCount int
	}{
		{
			name:             "valid image name",
			image:            "alpine",
			expectedMinCount: 10,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			provider := &dockerToolProvider{}
			request := mcp.CallToolRequest{}
			request.Params.Arguments = map[string]interface{}{
				"image": c.image,
			}

			result, err := provider.GetImageVersions(context.Background(), request)
			assert.NoError(t, err)

			textContent, ok := result.Content[0].(mcp.TextContent)
			assert.True(t, ok)

			versions := strings.Split(textContent.Text, "\n")
			assert.Greater(t, len(versions), c.expectedMinCount)
		})
	}
}
