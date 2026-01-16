# Backend (server)

The backend is a Go HTTP API.

## Tech Stack

- Go (standard library for HTTP, more libraries will be added later)
- Planned database: PostgreSQL

## Development

From the `server/` directory:

```bash
go run main.go
```

## The API will be available at:

- Health check: GET http://localhost:8080/health

## Endpoints (current)

- GET /health – basic health check returning `{"status":"ok"}`