# w3r5l3y-ssh

A personal terminal website served over SSH.

Built with Go, Wish, Bubble Tea, Lip Gloss, and Docker.

## Overview

`w3r5l3y-ssh` is a public-facing SSH application that presents a minimalist terminal interface instead of a traditional web page.

It includes:

* About page
* Projects page
* Source page
* Contact page
* Keyboard navigation
* Clickable terminal links in supported terminals
* Docker support

## Run locally

Start the SSH server:

```bash
go run ./cmd/server
```

Connect from another terminal:

```bash
ssh localhost -p 23234
```

## Run with Docker Compose

Build and start the container:

```bash
docker compose up --build
```

Connect from another terminal:

```bash
ssh localhost -p 23234
```

## Configuration

Environment variables:

| Variable            |                      Default | Description                      |
| ------------------- | ---------------------------: | -------------------------------- |
| `SSH_ADDR`          |              `0.0.0.0:23234` | Address and port for the SSH app |
| `SSH_HOST_KEY_PATH` | `/data/ssh_host_ed25519_key` | Persistent SSH host key path     |
| `SSH_THEME`         |                     `violet` | Colour theme                     |

## Docker image

The project is designed to be built as a container image and deployed as a self-hosted service.

Planned image:

```text
ghcr.io/w3r5l3y/w3r5l3y-ssh:latest
```

## Notes

This is a custom SSH application, not a normal shell login.

The SSH host key should be persisted between container restarts so clients do not see host identity warnings after updates.
