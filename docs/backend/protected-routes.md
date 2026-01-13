# Protected API Routes

All sensitive routes are protected with JWT middleware.

## Authorization

Clients must send:

Authorization: Bearer <token>

## Example

GET /api/vault/me
