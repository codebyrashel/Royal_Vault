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

## Data Model (ERD-Level)

### Entities

**User**
- `id` (UUID or integer)
- `email` (unique)
- `password_hash` (for login authentication)
- `created_at`
- `updated_at`

**Vault**
- `id`
- `user_id` (FK → User.id, 1:1)
- `encrypted_vault_key` (string or bytea)
- `created_at`
- `updated_at`

**Folder**
- `id`
- `user_id` (FK → User.id)
- `name`
- `created_at`
- `updated_at`

**Credential**
- `id`
- `vault_id` (FK → Vault.id)
- `folder_id` (FK → Folder.id, nullable)
- `title` (may be stored in plaintext)
- `url` (may be stored in plaintext)
- `encrypted_username`
- `encrypted_password`
- `encrypted_notes` (nullable)
- `created_at`
- `updated_at`

**SecurityQuestion**
- `id`
- `credential_id` (FK → Credential.id)
- `encrypted_question`
- `encrypted_answer`
- `created_at`
- `updated_at`

### Relationships

- User 1 — 1 Vault
- User 1 — N Folders
- Vault 1 — N Credentials
- Folder 1 — N Credentials (optional)
- Credential 1 — N SecurityQuestions