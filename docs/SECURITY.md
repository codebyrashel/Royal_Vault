# Security Model

This document outlines the high-level security model for Royal Vault.

## Zero-Knowledge Principle

- The server never sees decrypted secrets.
- All sensitive data (usernames, passwords, notes, security questions and answers) is encrypted on the client.
- The master password is never sent to the backend.

## Keys

There are two important password/keys concepts:

1. **Login Password**
   - Used only for authenticating the user with the backend.
   - Stored on the server as a password hash (e.g., bcrypt).

2. **Master Password**
   - Known only to the client.
   - Used to derive a cryptographic key (e.g., via PBKDF2 or Argon2).
   - That derived key encrypts/decrypts the vault key.

## Vault Key

- A random symmetric key (e.g., AES-GCM key) generated in the client.
- Used to encrypt:
  - Credential usernames
  - Credential passwords
  - Notes
  - Security questions and answers
- The vault key itself is encrypted with a key derived from the master password.
- The backend only stores:
  - The encrypted vault key
  - Encrypted credential fields