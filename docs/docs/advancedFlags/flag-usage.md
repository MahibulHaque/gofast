---
sidebar_position: 1
sidebar_label: Advanced Flag Usage
---

# Advaned Flags


The `--advanced` flag in gofast serves as a switch to enable additional features during project creation. It is applied with the create command and unlocks the following features:

- **CI/CD Workflow Setup using GitHub Actions**: Automates the setup of a CI/CD workflow using GitHub Actions.

- **Websocket Support**: WebSocket endpoint that sends continuous data streams through the WS protocol.

- **Docker**: Docker configuration for go project.

- **React**: Frontend written in TypeScript, including an example fetch request to the backend.

To utilize the `--advanced` flag, use the following command:

```bash
gofast create --name <project_name> --framework <selected_framework> --driver <selected_driver> --advanced
```

By including the `--advanced` flag, users can choose one or all of the advanced features. The flag enhances the simplicity of the cli while offering flexibility for users who require additional functionality.

To recreate the project using the same configuration semi-interactively, use the following command:
```bash
gofast create --name my-project --framework chi --driver mysql --advanced
```

Non-Interactive Setup is also possible:
```bash
gofast create --name my-project --framework chi --driver mysql --advanced --feature githubaction --feature websocket```
