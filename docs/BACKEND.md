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