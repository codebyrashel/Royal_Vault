# Architecture Overview

This document describes the overall system design of Royal Vault, a secure credential manager.

## High-Level Goals

- Zero-knowledge: the server never sees decrypted secrets.
- Clear separation of concerns:
  - `client/` – UI and cryptography (encryption/decryption in the browser).
  - `server/` – authentication, API endpoints, and persistence.
  - PostgreSQL – structured storage for users, vault metadata, and encrypted data.

## Current State

```markdown
- Frontend:
  - Vue 3 + TypeScript app in `client/`, created with Vite.
- Backend:
  - Go HTTP server in `server/` with a basic `/health` endpoint.
```
These components are currently independent. API endpoints and data models will be added in later steps.

### UI State

- Routing is set up with Vue Router:
  - Landing page (`/`)
  - Dashboard placeholder (`/dashboard`)
- No authentication or API integration yet.

More details will be added as the project evolves.