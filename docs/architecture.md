# Royal Vault Architecture

Royal Vault is a zero-knowledge password manager.

## System Design

- Vue frontend performs all encryption
- Golang API stores encrypted vault blobs
- PostgreSQL persists encrypted data
- Server never sees plaintext