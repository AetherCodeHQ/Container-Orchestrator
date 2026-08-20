# Container Orchestrator

![CI](https://github.com/Qyroxen/Container-Orchestrator/actions/workflows/ci.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go) ![License](https://img.shields.io/badge/License-MIT-yellow.svg) ![Stars](https://img.shields.io/github/stars/Qyroxen/Container-Orchestrator?style=social)

> Simplified container orchestration - manage Docker like a pro

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Container-Orchestrator?style=social)](https://github.com/Qyroxen/Container-Orchestrator/stargazers)

## What is it?

Container Orchestrator provides a simple CLI for managing Docker containers, networks, and volumes with powerful features.

## Why should you care?

Docker commands are verbose. This tool simplifies container management.

## Demo

```bash
./container-orch ps
```

**Output:**
```
CONTAINER ID   IMAGE          STATUS      PORTS
a1b2c3d4e5f6   nginx:latest   running     0.0.0.0:80->80/tcp
b2c3d4e5f6g7   postgres:13    running     0.0.0.0:5432->5432/tcp
```

## Features

- Container management
- Network management
- Volume management
- Docker Compose support
- Health monitoring

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Container-Orchestrator.git
cd Container-Orchestrator
go build -o container-orch .

# Run
./container-orch ps
```

## CLI Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--format` | Output format (table, json) | `table` |
| `--filter` | Filter containers | `none` |
| `--watch` | Real-time monitoring | `false` |

## Examples

# List containers
./container-orch ps

# Start containers
./container-orch up --detach

# View logs
./container-orch logs --follow

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Container-Orchestrator/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Container-Orchestrator?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Container-Orchestrator/network/members">
    <img src="https://img.shields.io/github/forks/Qyroxen/Container-Orchestrator?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Container-Orchestrator/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Container-Orchestrator" alt="Issues">
  </a>
</p>
