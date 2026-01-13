# JWT Authentication

Royal Vault uses JWT tokens for session authentication.

## Flow

1. User logs in with email + password
2. Server verifies password hash
3. Server issues JWT token
4. Client sends token in Authorization header

## Security

- Token signed with HS256
- Contains user ID and email
- Stateless authentication
