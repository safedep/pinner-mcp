# Contributing to Pinner MCP

## Development Setup

### Prerequisites

- Docker
- Go

### Docker Image Configuration

In the `.cursor/mcp.json` file, you'll notice that the Docker image is set to `safedep/pinner-mcp:latest` instead of `ghcr.io/safedep/pinner-mcp:latest`. This configuration is specifically designed for local development and testing purposes.

The `safedep/pinner-mcp:latest` tag allows you to build and test the Docker image locally using:

```bash
docker build -t safedep/pinner-mcp:latest .
```

This local image tag matches the configuration in `.cursor/mcp.json`, enabling seamless integration with Cursor IDE for local development. The `ghcr.io/safedep/pinner-mcp:latest` tag, on the other hand, is used for the official releases hosted on GitHub Container Registry.

## How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Build and test locally using the Docker configuration mentioned above
5. Commit your changes (`git commit -m 'Add some amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request
