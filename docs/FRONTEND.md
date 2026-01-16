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