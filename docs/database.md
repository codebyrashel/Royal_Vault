# Database Design

Database: PostgreSQL (Dockerized)

## Users Table

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    salt TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```
### Security
- Raw passwords are never stored
- Passwords are hashed with bcrypt
- Unique salt per user