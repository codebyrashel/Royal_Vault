# Backend (server)

The backend is a Go HTTP API.

## Tech Stack

- Go
- Gin (HTTP routing)
- Planned database: PostgreSQL

## Database

- PostgreSQL is used as the primary database.
- For local development, PostgreSQL typically runs in a Docker container:
  - Container name: `royal_vault_db` (suggested)
  - Host: `localhost`
  - Port: `5433` (example)
  - User: `ROYAL_VAULT_DB_USER` (from local env config)
  - Password: `ROYAL_VAULT_DB_PASSWORD` (from local env config)
  - Database: `ROYAL_VAULT_DB_NAME` (from local env config)
- Connection details are read from environment variables:
  - `DB_HOST` (e.g., `localhost`)
  - `DB_PORT` (e.g., `5433`)
  - `DB_USER`
  - `DB_PASSWORD`
  - `DB_NAME`
  - `DB_SSLMODE` (typically `disable` for local dev)

- `vaults` table stores:
  - `user_id` (1:1 with users)
  - `encrypted_vault_key` (opaque string from the client)
  - `salt` (base64-encoded salt used for deriving the master key)

### Environment Configuration

Local environment variables are stored in a `.env` file (not committed to Git). An example template is provided in `.env.example`.

Before running the backend locally, export the variables from `.env` into your shell:

```bash
export $(grep -v '^#' .env | xargs)
cd server
go run main.go
```
- Alternatively, you can configure these variables in your shell or development environment.

## Project Structure (current)

- `server/main.go` – application entry point and server startup
- `server/internal/handlers` – HTTP handlers (e.g., health check)
- `server/internal/routes` – route registration and Gin engine setup
- `server/internal/db` – database connection logic (PostgreSQL via database/sql and lib/pq)

## Development

From the `server/` directory:

```bash
go run main.go
```

## The API will be available at:

- Health check: GET http://localhost:8080/health

## Endpoints (current)

- `GET /health` – basic health check returning `{"status":"ok"}`
- `POST /auth/signup` – creates a new user and associated vault (stores `encryptedVaultKey` and `salt`)
- `POST /auth/login` – authenticates a user and returns a JWT, `encryptedVaultKey`, and `salt`
- `GET /credentials` – list all credentials for the authenticated user (encrypted fields).
- `POST /credentials` – create a new credential (expects encrypted fields from the client).

> Note that endpoints are protected with Bearer JWT.

## CORS

For local development, the backend enables CORS for the frontend origin:

- Allowed origin: `http://localhost:5173`
- Allowed methods: `GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`
- Allowed headers: `Origin`, `Content-Type`, `Authorization`

This is configured via `github.com/gin-contrib/cors` in `server/internal/routes/routes.go`.

## Planned API Design

All endpoints (except signup/login) require authentication.

### Auth

- `POST /auth/signup`
  - Request body:
    - `email`
    - `password` (login password)
    - `encryptedVaultKey` (string)
  - Response:
    - `userId`
    - `token` (JWT or similar)

- `POST /auth/login`
  - Request body:
    - `email`
    - `password` (login password)
  - Response:
    - `userId`
    - `token`
    - `encryptedVaultKey`

### Vault

- `GET /vault`
  - Returns:
    - Vault metadata (if needed)
    - `encryptedVaultKey` (may be returned here or via login)

### Folders

- `GET /folders`
- `POST /folders`
  - Body:
    - `name`
- `PUT /folders/:id`
  - Body:
    - `name`
- `DELETE /folders/:id`

### Credentials

- `GET /credentials`
  - Returns a list of credentials for the authenticated user.
- `GET /credentials/:id`
- `POST /credentials`
  - Body:
    - `title`
    - `url`
    - `folderId` (optional)
    - `encryptedUsername`
    - `encryptedPassword`
    - `encryptedNotes` (optional)
- `PUT /credentials/:id`
- `DELETE /credentials/:id`

### Security Questions

- `GET /credentials/:credentialId/security-questions`
- `POST /security-questions`
  - Body:
    - `credentialId`
    - `encryptedQuestion`
    - `encryptedAnswer`
- `PUT /security-questions/:id`
- `DELETE /security-questions/:id`