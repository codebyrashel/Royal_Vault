# Project Setup Guide

## Requirements

- Go 1.22+
- Docker & Docker Compose
- Node.js (for frontend later)
- PostgreSQL client (psql)

### Clone Repository

```bash
git clone <repo-url>
cd Royal_Vault

# Start Database

sudo docker compose up -d

# Run Backend

cd server
go run main.go

# Server runs at:
https://localhost:5000
```