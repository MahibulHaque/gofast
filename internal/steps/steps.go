package steps

import "github.com/mahibulhaque/gofast/internal/flags"

type StepSchema struct {
	StepName string
	Options  []Item
	Headers  string
	Field    string
}

type Steps struct {
	Steps map[string]StepSchema
}

type Item struct {
	Flag, Title, Desc string
}

func InitSteps(projectType flags.Framework, databaseType flags.Database) *Steps {
	steps := &Steps{
		map[string]StepSchema{
			"framework": {
				StepName: "Go Project Framework",
				Options: []Item{
					{
						Title: "Standard-Library",
						Desc:  "The bulit-in Go standard library",
					},
					{
						Title: "Chi",
						Desc:  "A lightweight, idiomatic and composable router for building Go HTTP services",
					},
					{
						Title: "Gin",
						Desc:  "Features a martini-like API with performance that is up to 40 times faster thanks to httprouter",
					},
					{
						Title: "Fiber",
						Desc:  "An Express inspired web framework built on top of Fasthttp",
					},
					{
						Title: "Gorilla/Mux",
						Desc:  "Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler",
					},
					{
						Title: "HttpRouter",
						Desc:  "HttpRouter is a lightweight high performance HTTP request router for Go",
					},
					{
						Title: "Echo",
						Desc:  "High performance, extensible, minimalist Go web framework",
					},
				},
				Headers: "What framework do you want to use in your Go project?",
				Field:   projectType.String(),
			},
			"driver": {
				StepName: "Go Project Database Driver",
				Options: []Item{
					{
						Title: "Mysql",
						Desc:  "MySQL-Driver for Go's database/sql package",
					},
					{
						Title: "Postgres",
						Desc:  "Go postgres driver for Go's database/sql package"},
					{
						Title: "Sqlite",
						Desc:  "sqlite3 driver conforming to the built-in database/sql interface"},
					{
						Title: "Mongo",
						Desc:  "The MongoDB supported driver for Go."},
					{
						Title: "Redis",
						Desc:  "Redis driver for Go."},
					{
						Title: "None",
						Desc:  "Choose this option if you don't wish to install a specific database driver."},
				},
				Headers: "What database driver do you want to use in your Go project?",
				Field:   databaseType.String(),
			},
			"advanced": {
				StepName: "Advanced Features",
				Headers:  "Which advanced features do you want?",
				Options: []Item{
					{
						Flag:  "React",
						Title: "React",
						Desc:  "Use Vite to spin up a React project in TypeScript.",
					},
					{
						Flag:  "GitHubAction",
						Title: "Go Project Workflow",
						Desc:  "Workflow templates for testing, cross-compiling and releasing Go projects",
					},
					{
						Flag:  "Websocket",
						Title: "Websocket endpoint",
						Desc:  "Add a websocket endpoint",
					},
					{
						Flag:  "Docker",
						Title: "Docker",
						Desc:  "Dockerfile and docker-compose generic configuration for go project",
					},
				},
			},
			"git": {
				StepName: "Git Repository",
				Headers:  "Which git option would you like to select for your project?",
				Options: []Item{
					{
						Title: "Commit",
						Desc:  "Initialize a new git repository and commit all the changes",
					},
					{
						Title: "Stage",
						Desc:  "Initialize a new git repository but only stage the changes",
					},
					{
						Title: "Skip",
						Desc:  "Proceed without initializing a git repository",
					},
				},
			},
		},
	}
	return steps
}
