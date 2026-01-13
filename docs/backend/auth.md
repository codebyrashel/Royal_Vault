# Authentication System

## Registration Flow

1. Client sends email and password
2. Server generates cryptographic salt
3. Password + salt is hashed with bcrypt
4. Hash and salt stored in database

## Security Guarantees

- No plaintext passwords
- Unique salt per user
- Resistant to rainbow-table attacks