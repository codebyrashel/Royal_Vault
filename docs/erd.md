# Database Design

## users
- id
- email
- password_hash
- salt
- created_at

## vaults
- id
- user_id
- encrypted_data
- updated_at