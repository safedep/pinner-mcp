# Pinner MCP 📍

A Model Context Protocol (MCP) server that can help pin GitHub Actions to a specific commit hash.

![Pinner MCP](./docs/assets/demo.png)

## Usage

Run as a container with `stdio` transport.

```bash
docker run -it --rm ghcr.io/safedep/pinner-mcp:latest
```

If you are using Cursor, you can add the following to your `.cursor/mcp.json` file.

```json
{
  "mcpServers": {
    "pinner-mcp-stdio-server": {
      "command": "docker",
      "args": [
        "run",
        "--rm",
        "-i",
        "safedep/pinner-mcp:latest"
      ]
    }
  }
}
```


