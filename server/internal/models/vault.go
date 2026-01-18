package models

import "time"

type Vault struct {
	ID               int64     `db:"id"`
	UserID           int64     `db:"user_id"`
	EncryptedVaultKey string   `db:"encrypted_vault_key"`
	Salt             string    `db:"salt"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}