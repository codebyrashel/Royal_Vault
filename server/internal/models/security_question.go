package models

import "time"

type SecurityQuestion struct {
	ID                int64     `db:"id"`
	CredentialID      int64     `db:"credential_id"`
	EncryptedQuestion string    `db:"encrypted_question"`
	EncryptedAnswer   string    `db:"encrypted_answer"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}