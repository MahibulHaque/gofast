# Gofast

<h2>
    Introducing the Ultimate Golang Starter Library
</h2>

A simple golang library, now available in your favourite terminal. Start new golang project from structured workflows of your choice.

## Why should you use this?

- Easy to set up and install
- Have the entire Go structure already established
- Setting up a Go HTTP server (or Fasthttp with Fiber)
- Integrate with a popular frameworks
- Focus on the actual code of your application

<a id="install"></a>

<h2>
  Installation
</h2>

```bash
go install github.com/mahibulhaque/gofast@latest
```

This installs a go binary that will automatically bind to your $GOPATH

> if you’re using Zsh, you’ll need to add it manually to `~/.zshrc`.

```bash
GOPATH=$HOME/go  PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

don't forget to update

```bash
source ~/.zshrc
```

Then in a new terminal run:

```bash
gofast create
```

You can also use the provided flags to set up a project without interacting with the UI.

```bash
gofast create --name myproject --framework chi --driver postgres --git skip
```

See `gofast create -h` for all the options and shorthands cli flags.

<a id="frameworks"></a>

<h2>
  Frameworks
</h2>

- [Chi](https://github.com/go-chi/chi)
- [Gin](https://github.com/gin-gonic/gin)
- [Fiber](https://github.com/gofiber/fiber)
- [HttpRouter](https://github.com/julienschmidt/httprouter)
- [Gorilla/mux](https://github.com/gorilla/mux)
- [Echo](https://github.com/labstack/echo)

<a id="database"></a>

<h2>
  Database
</h2>

Gofast offers enhanced database support, allowing you to choose your preferred database driver during project setup. Use the `--driver` or `-d` flag to specify the database driver you want to integrate into your project.

### Supported Database Drivers

Choose from a variety of supported database drivers:

- [Mongo](https://go.mongodb.org/mongo-driver)
- [Mysql](https://github.com/go-sql-driver/mysql)
- [Postgres](https://github.com/jackc/pgx/)
- [Sqlite](https://github.com/mattn/go-sqlite3)
- [Redis](https://github.com/redis/go-redis)

<a id="advanced-features"></a>

<h2>
  Advanced Features
</h2>

The tool is focused on being as minimalistic as possible. That being said, we wanted to offer the ability to add other features people may want without bloating the overall experience.

You can now use the `--advanced` flag when running the `create` command to get access to the following features. This is a multi-option prompt; one or more features can be used at the same time:

- CI/CD workflow setup using [Github Actions](https://docs.github.com/en/actions)
- [Websocket](https://pkg.go.dev/github.com/coder/websocket) sets up a websocket endpoint
- Docker configuration for go project
- [React](https://react.dev/) frontend written in TypeScript, including integration with [Tanstack Router](https://tanstack.com/router/latest) and [Tanstack Query](https://tanstack.com/query/latest)

<a id="usage"></a>

<h2>
  Usage
</h2>

Here's an example of setting up a project with a specific database driver:

```bash
gofast create --name myproject --framework chi --driver postgres --git commit
```
