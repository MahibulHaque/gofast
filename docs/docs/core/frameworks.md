---
sidebar_position: 1
sidebar_label: Golang Frameworks
---

# Frameworks

Created projects can utilize several Go web frameworks to handle HTTP routing and server functionality. The chosen frameworks are:

- [Chi](https://github.com/go-chi/chi): Lightweight and flexible router for building Go HTTP services.
- [Echo](https://github.com/labstack/echo): High-performance, extensible, minimalist Go web framework.
- [Fiber](https://github.com/gofiber/fiber): Express-inspired web framework designed to be fast, simple, and efficient.
- [Gin](https://github.com/gin-gonic/gin): A web framework with a martini-like API, but with much better performance.
- [Gorilla/Mux](https://github.com/gorilla/mux): A powerful URL router and dispatcher for Golang.
- [HttpRouter](https://github.com/julienschmidt/httprouter): A high-performance HTTP request router that scales well.

## Standard Project Structure

The project is structured with a simple layout, focusing on the cmd, internal directories, allowing high *Test driven development* within project:

```bash
└── (Root)/
    ├── cmd/
    │   └── api/
    │       └── main.go
    ├── internal/
    │   ├── db/
    │   │   ├── database.go
    │   │   └── database_test.go
    │   ├── migrations/
    │   │   └── 0001_create_user_table.sql
    │   ├── repository/
    │   │   ├── auth.go
    │   │   ├── user.go
    │   │   └── post.go
    │   ├── server/
    │   │   ├── server.go
    │   │   └── routes.go
    │   ├── middlewares/
    │   │   ├── logger.go
    │   │   └── routeGuard.go
    │   ├── handlers/
    │   │   ├── auth.go
    │   │   ├── auth_test.go
    │   │   ├── user.go
    │   │   ├── user_test.go
    │   │   ├── post.go
    │   │   └── post_test.go
    │   └── services/
    │       ├── auth.go
    │       ├── auth_test.go
    │       ├── user.go
    │       ├── user_test.go
    │       ├── post.go
    │       └── post_test.go
    ├── go.mod
    ├── go.sum
    ├── Makefile
    └── README.md
```

