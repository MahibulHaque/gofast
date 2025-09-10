---
sidebar_position: 1
sidebar_label: Project Init
---

# Creating a Project

After installing the Gofast CLI tool, you can create a new project with the default settings by running the following command:

```bash
gofast create
```

This command will interactively guide you through the project setup process, allowing you to choose the project name, framework, and database driver.

### Using Flags for Non-Interactive Setup

For a non-interactive setup, you can use flags to provide the necessary information during project creation. Here's an example:

```bash
gofast create --name new-project --framework chi --driver mysql --git commit
```

In this example:

- `--name`: Specifies the name of the project (replace "new-project" with your desired project name).
- `--framework`: Specifies the Go framework to be used (e.g., "chi").
- `--driver`: Specifies the database driver to be integrated (e.g., "mysql").
- `--git`: Specifies the git configuration option of the project (e.g., "commit").
Customize the flags according to your project requirements.

### Advanced Flag

By including the `--advanced` flag, users can choose one or all of the advanced features, GitHub Actions for CI/CD, Websocket, Docker, during the project creation process. The flag enhances the simplicity of the cli while offering flexibility for users who require additional functionality.

```bash
gofast create --advanced
```

To recreate the project using the same configuration semi-interactively, use the following command:

```bash
gofast create --name my-project --framework chi --driver mysql --git commit --advanced
```


This approach opens interactive mode only for advanced features, which allows you to choose the one or combination of available features.

