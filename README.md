# Pinner MCP 📍

A Model Context Protocol (MCP) server that can help pin 3rd party dependencies to immutable digests.
Supported dependency types include:

- Docker base images
- GitHub Actions

![Pinner MCP](./docs/assets/demo.png)

## 📦 Usage

Run as a container with `stdio` transport.

```bash
docker run -it --rm ghcr.io/safedep/pinner-mcp:latest
```

### 💻 VS Code

Add the following to your `.vscode/mcp.json` file in your workspace. You must have the GitHub Copilot extension installed and enabled.

```json
{
  "servers": {
    "pinner-mcp": {
      "type": "stdio",
      "command": "docker",
      "args": ["run", "--rm", "-i", "ghcr.io/safedep/pinner-mcp:latest"]
    }
  }
}
```

Use GitHub Copilot Chat with prompts like:

```
Pin GitHub Actions to their commit hash
```

```
Pin container base images to digests
```

To update pinned versions, you can use a prompt like the following:

```
Update pinned versions of container base images
```

### 💻 Cursor

Add the following to your `.cursor/mcp.json` file. You must _enable_
the MCP server in the settings. Learn more [here](https://docs.cursor.com/context/model-context-protocol#what-is-mcp).

```json
{
  "mcpServers": {
    "pinner-mcp-stdio-server": {
      "command": "docker",
      "args": ["run", "--rm", "-i", "ghcr.io/safedep/pinner-mcp:latest"]
    }
  }
}
```

Use a Composer prompt like the following to pin a specific commit hash.

```
Pin GitHub Actions to their commit hash
```

```
Pin container base images to digests
```

To update pinned versions, you can use a prompt like the following.

```
Update pinned versions of container base images
```

### 🔄 Tool Updates

Updates for the MCP server are automatically pushed to the `latest` tag on
[GitHub Container Registry](https://github.com/safedep/pinner-mcp/pkgs/container/pinner-mcp). You
must manually update your local container image to the latest version.

```bash
docker pull ghcr.io/safedep/pinner-mcp:latest
```

## 🔧 Development

### Building Locally

For local development on this project, build the container image locally:

```bash
docker build -t pinner-mcp:local .
```

The `.vscode/mcp.json` and `.cursor/mcp.json` files in this repository are configured to use the local image (`pinner-mcp:local`) for development and testing of unpublished changes.

After building locally, you can use VS Code or Cursor with GitHub Copilot as described in the usage sections above.

## 📚 References

- Originally built to protect [vet](https://github.com/safedep/vet) from malicious GitHub Actions
- [mcp-go](https://github.com/mark3labs/mcp-go) is a great library for building MCP servers
- Built and maintained by [SafeDep Engineering](https://safedep.io)
