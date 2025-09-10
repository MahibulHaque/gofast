---
sidebar_position: 5
sidebar_label: React & Tanstack (Vite)
---

# React & Tanstack

This template provides a minimal setup for getting React working with Vite and Tanstack for the frontend. It comes out of the box with integration of React with [Tanstack Router](https://tanstack.com/router/latest), [Tanstack Query](https://tanstack.com/query/latest) and [TailwindCSS v4](https://tailwindcss.com/) for rapid type-safe development.

## React Project Structure Overview
```bash
(Root)
└── frontend/
    ├── node_modules/
    ├── public/
    │   ├── index.html
    │   └── favicon.ico
    ├── src/
    │   ├── components/
    │   │   └── Header.tsx
    │   ├── lib/
    │   │   └── utils.ts
    │   ├── routes/
    │   │   ├── __root.tsx  
    │   │   ├── index.tsx
    │   │   └── demo.tanstack-query.tsx
    │   ├── styles.css
    │   └── main.tsx
    ├── .env
    ├── components.json
    ├── prettier.config.js
    ├── eslint.config.js
    ├── index.html
    ├── package.json
    ├── package-lock.json
    ├── vite.config.ts
    ├── tsconfig.app.json
    ├── tsconfig.json
    ├── tsconfig.node.json
    └── README.md
```

## Usage
1. **Navigate to the `frontend` directory**: First, navigate to the `frontend` directory where the React project resides.
```bash
cd frontend
```
2. **Install Dependencies**: Use npm to install all necessary dependencies.
```bash
npm install
```
3. **Run the Development Server**: Start the Vite development server for local development. This will launch a live-reloading server on a default port.
```bash
npm run dev
```

You should now be able to access the React application by opening a browser and navigating to `http://localhost:5173`.

## Makefile

The make run target will start the Go server in the backend, install frontend dependencies, and run the Vite development server for the frontend.

```make
run:
    @go run cmd/api/main.go &
    @npm install --prefix ./frontend
    @npm run dev --prefix ./frontend
```

## Dockerfile

Combine React advanced flag with Docker flag to get Docker and docker-compose configuration and run them with:

```bash
make docker-run
```

### Dockerfile with frontend

```dockerfile
FROM golang:1.25.0-alpine AS build

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

### Docker compose without database
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
  frontend:
    build:
      context: .
      dockerfile: Dockerfile
      target: frontend
    restart: unless-stopped
    ports:
      - 5173:5173
    depends_on:
      - app
```

### Docker compose with database

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
      DB_SCHEMA: ${DB_SCHEMA}
    depends_on:
      psql:
        condition: service_healthy
    networks:
      - gofast
  frontend:
    build:
      context: .
      dockerfile: Dockerfile
      target: frontend
    restart: unless-stopped
    depends_on:
      - app
    ports:
      - 5173:5173
    networks:
      - gofast
  psql:
    image: postgres:latest
    restart: unless-stopped
    environment:
      POSTGRES_DB: ${DB_DATABASE}
      POSTGRES_USER: ${DB_USERNAME}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
    ports:
      - "${DB_PORT}:5432"
    volumes:
      - psql_volume:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "sh -c 'pg_isready -U ${DB_USERNAME} -d ${DB_DATABASE}'"]
      interval: 5s
      timeout: 5s
      retries: 3
      start_period: 15s
    networks:
      - gofast

volumes:
  psql_volume:
networks:
  gofast:
```

## Environment Variables

The `VITE_PORT` in `.env` refers `PORT` from `.env` in project root ( for backend ). If value of `PORT` is changed than `VITE_PORT` must also be changed so that requests to backend work fine and have no conflicts.


