# Frontend (client)

The frontend is a Vue 3 application built with Vite and TypeScript.

## Tech Stack

- Vue 3 (Composition API)
- TypeScript
- Vite

## Development

From the `client/` directory:

```bash
npm install
npm run dev
```
The app will be available at `http://localhost:5173` by default.

## Current Pages

- `/` – Landing page with:
  - Hero section
  - How it works
  - Security overview (high-level)
  - FAQ placeholder
- `/dashboard` – Placeholder dashboard page that will later show the user's vault and credentials.

## Planned User Flows

### Signup Flow

1. User visits `/signup` (to be implemented).
2. User enters:
   - Email
   - Login password
   - Master password
3. Client:
   - Derives a key from the master password.
   - Generates a random vault key.
   - Encrypts the vault key with the derived key to produce `encryptedVaultKey`.
4. Client sends to backend:
   - Email
   - Login password
   - `encryptedVaultKey`

### Login Flow

1. User visits `/login` (to be implemented).
2. User enters:
   - Email
   - Login password
   - Master password
3. Backend validates login password and returns:
   - Auth token
   - `encryptedVaultKey`
4. Client:
   - Derives key from master password.
   - Decrypts `encryptedVaultKey` to get the vault key.
   - Uses vault key to decrypt credentials retrieved from the API.

## Cryptography Utilities

The frontend implements client-side encryption helpers in `src/utils/crypto.ts`:

- `deriveKeyFromPassword(password, salt?)`  
  Derives a symmetric key from a master password using PBKDF2 (`SHA-256`, 150,000 iterations). Returns:
  - `key`: `CryptoKey` for AES-GCM encryption/decryption
  - `salt`: `Uint8Array` used for derivation (must be stored and reused on login)

- `generateVaultKey()`  
  Generates a random 256-bit AES-GCM key used as the vault key.

- `exportKey(key)` / `importVaultKey(raw)`  
  Convert between `CryptoKey` and raw `Uint8Array` form for storage/encryption.

- `encryptText(key, plaintext)` / `decryptText(key, payload)`  
  Encrypt/decrypt arbitrary strings using AES-GCM. IV and ciphertext are returned/stored as base64 strings.

There is a temporary development page at `/crypto-test` that verifies:

- A vault key can be encrypted and decrypted using a master password-derived key.
- Arbitrary plaintext can be encrypted and decrypted correctly using the vault key.


## Current Auth UI

- `/signup`
  - Accepts:
    - Email
    - Login password
    - Master password
  - Client:
    - Derives a master key from the master password.
    - Generates a random vault key.
    - Encrypts the vault key with the master key.
    - Sends `email`, `login password`, `encryptedVaultKey`, and `salt` to the backend.

- `/login`
  - Accepts:
    - Email
    - Login password
    - Master password
  - Client:
    - Calls `/auth/login` (email + login password).
    - Receives `token`, `encryptedVaultKey`, and `salt`.
    - Re-derives the master key using the master password + salt.
    - Decrypts the vault key.
    - Stores JWT token and decrypted vault key in an in-memory store.
    - Redirects to `/dashboard`.