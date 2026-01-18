package models

import "time"

type Credential struct {
	ID                int64     `db:"id"`
	VaultID           int64     `db:"vault_id"`
	FolderID          *int64    `db:"folder_id"`
	Title             string    `db:"title"`
	URL               *string   `db:"url"`
	EncryptedUsername string    `db:"encrypted_username"`
	EncryptedPassword string    `db:"encrypted_password"`
	EncryptedNotes    *string   `db:"encrypted_notes"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}