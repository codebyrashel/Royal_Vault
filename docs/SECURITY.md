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

## Encryption Flow (Client-Side)

### Key Concepts

- **Master Password**
  - Entered by the user in the client.
  - Never sent to the backend.
  - Used to derive a `master_key` via a KDF (e.g., PBKDF2).

- **Master Key**
  - Derived from `master_password + salt`.
  - Used only to encrypt/decrypt the vault key.

- **Vault Key**
  - Random symmetric key (e.g., 256-bit).
  - Used to encrypt/decrypt all sensitive fields:
    - `encrypted_username`
    - `encrypted_password`
    - `encrypted_notes`
    - `encrypted_question`
    - `encrypted_answer`

### Signup Flow (Encryption Perspective)

1. User enters:
   - Email
   - Login password
   - Master password
2. Client:
   - Generates a random salt for the master key derivation.
   - Derives `master_key` from `master_password + salt` (PBKDF2).
   - Generates a random `vault_key`.
   - Encrypts `vault_key` with `master_key` → `encrypted_vault_key`.
3. Client sends to backend:
   - Email
   - Login password
   - `encrypted_vault_key`
   - (Optionally the salt used for master key derivation, if stored server-side)
4. Backend stores:
   - `users.password_hash`
   - `vaults.encrypted_vault_key`
   - (Optionally master key derivation salt)

### Login Flow (Encryption Perspective)

1. User enters:
   - Email
   - Login password
   - Master password
2. Backend verifies login and returns:
   - JWT token
   - `encrypted_vault_key`
   - (Optionally master key derivation salt)
3. Client:
   - Re-derives `master_key` from `master_password + salt`.
   - Decrypts `encrypted_vault_key` → `vault_key`.
   - Uses `vault_key` to decrypt any encrypted credential data fetched from the API.

## Client-Side Encryption Implementation

- Key derivation:
  - Implemented in `client/src/utils/crypto.ts` using PBKDF2 (`SHA-256`, 150k iterations).
  - Inputs: master password + random salt.
  - Output: `master_key` (AES-GCM key) and the salt.
  - The salt must be stored (e.g., with the vault) so the same master key can be derived on login.

- Vault key:
  - Generated randomly on signup (`generateVaultKey()`).
  - Exported to raw bytes and then encrypted with the `master_key`.
  - Only the encrypted vault key is sent to the backend.

- Data encryption:
  - All sensitive fields (usernames, passwords, notes, security questions/answers) are encrypted with the vault key using AES-GCM.
  - Encrypted payloads are stored as base64-encoded IV + ciphertext strings and treated as opaque on the server.