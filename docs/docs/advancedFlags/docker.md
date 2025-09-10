---
sidebar_position: 4
sidebar_label: Docker
---

# Docker

The Docker advanced flag provides the app's Dockerfile configuration and creates or updates the docker-compose.yml file, which is generated if a DB driver is used. The Dockerfile includes a two-stage build, and the final config depends on the use of advanced features. In the end, you will have a smaller image without unnecessary build dependencies.

## Dockerfile

```dockerfile
FROM golang:1.24.5-alpine AS build

RUN apk add --no-cache curl libstdc++ libgcc

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]
```

Docker config if React flag is used:
```dockerfile
FROM golang:1.24.5-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]


FROM node:20 AS frontend_builder
WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install
COPY frontend/. .
RUN npm run build

FROM node:20-slim AS frontend
RUN npm install -g serve
COPY --from=frontend_builder /frontend/dist /app/dist
EXPOSE 5173
CMD ["serve", "-s", "/app/dist", "-l", "5173"]
```

## Docker Compose

Docker and `docker-compose.yml` pull environment variables from the .env file.

Example if the Docker flag is used with the MySQL DB driver:
```yml
services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
      target: prod
    restart: unless-stopped
    ports:
      - ${PORT}:${PORT}
    environment:
      APP_ENV: ${APP_ENV}
      PORT: ${PORT}
      DB_HOST: ${DB_HOST}
      DB_PORT: ${DB_PORT}
      DB_DATABASE: ${DB_DATABASE}
      DB_USERNAME: ${DB_USERNAME}
      DB_PASSWORD: ${DB_PASSWORD}
    depends_on:
      mysql:
        condition: service_healthy
    networks:
      - gofast
  mysql:
    image: mysql:latest
    restart: unless-stopped
    environment:
      MYSQL_DATABASE: ${DB_DATABASE}
      MYSQL_USER: ${DB_USERNAME}
      MYSQL_PASSWORD: ${DB_PASSWORD}
      MYSQL_ROOT_PASSWORD: ${DB_ROOT_PASSWORD}
    ports:
      - "${DB_PORT}:3306"
    volumes:
      - mysql_volume:/var/lib/mysql
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "${DB_HOST}", "-u", "${DB_USERNAME}", "--password=${DB_PASSWORD}"]
      interval: 5s
      timeout: 5s
      retries: 3
      start_period: 15s
    networks:
      - gofast

volumes:
  mysql_volume:
networks:
  gofast:
```


