# Architecture Overview

This document describes the overall system design of Royal Vault, a secure credential manager.

## High-Level Goals

- Zero-knowledge: the server never sees decrypted secrets.
- Clear separation of concerns:
  - `client/` – UI and cryptography (encryption/decryption in the browser).
  - `server/` – authentication, API endpoints, and persistence.
  - PostgreSQL – structured storage for users, vault metadata, and encrypted data.

More details will be added as the project evolves.