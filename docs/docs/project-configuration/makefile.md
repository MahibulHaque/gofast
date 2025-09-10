---
sidebar_position: 2
---

# Makefile

### Makefile for project management

Makefile is designed for building, running, and testing a Go project. It handles OS-specific operations for Unix-based systems (Linux/macOS) and Windows.

### Targets

`all`

The default target that builds and test the application by running the `build` and `test` target.

`build`

Builds the Go application

`run`

Runs the Go application by executing the cmd/api/main.go file.

`docker-run` and `docker-down`

These targets manage a database container:

- **Unix-based systems**: Tries Docker Compose V2 first, falls back to V1 if needed.
- **Windows**: Uses Docker Compose without version fallback.

`clean`

Removes the compiled binary (`main` or `main.exe` depending on the OS).

`watch`

Enables live reload for the project using the `air` tool:

- **Unix-based systems**: Checks if air is installed and prompts for installation if missing.
- **Windows**: Uses PowerShell to manage air installation and execution.
