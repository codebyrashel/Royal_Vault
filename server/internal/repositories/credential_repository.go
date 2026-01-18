package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codebyrashel/Royal_Vault/server/internal/models"
)

type CredentialRepository struct {
	db *sql.DB
}

func NewCredentialRepository(db *sql.DB) *CredentialRepository {
	return &CredentialRepository{db: db}
}

func (r *CredentialRepository) CreateCredential(
	ctx context.Context,
	vaultID int64,
	folderID *int64,
	title string,
	url *string,
	encryptedUsername string,
	encryptedPassword string,
	encryptedNotes *string,
) (*models.Credential, error) {
	query := `
		INSERT INTO credentials (
			vault_id,
			folder_id,
			title,
			url,
			encrypted_username,
			encrypted_password,
			encrypted_notes,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING
			id, vault_id, folder_id, title, url,
			encrypted_username, encrypted_password, encrypted_notes,
			created_at, updated_at
	`

	row := r.db.QueryRowContext(
		ctx,
		query,
		vaultID,
		folderID,
		title,
		url,
		encryptedUsername,
		encryptedPassword,
		encryptedNotes,
	)

	var c models.Credential
	if err := row.Scan(
		&c.ID,
		&c.VaultID,
		&c.FolderID,
		&c.Title,
		&c.URL,
		&c.EncryptedUsername,
		&c.EncryptedPassword,
		&c.EncryptedNotes,
		&c.CreatedAt,
		&c.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to insert credential: %w", err)
	}

	return &c, nil
}

func (r *CredentialRepository) ListCredentialsByVaultID(
	ctx context.Context,
	vaultID int64,
) ([]*models.Credential, error) {
	query := `
		SELECT
			id, vault_id, folder_id, title, url,
			encrypted_username, encrypted_password, encrypted_notes,
			created_at, updated_at
		FROM credentials
		WHERE vault_id = $1
		ORDER BY title ASC
	`

	rows, err := r.db.QueryContext(ctx, query, vaultID)
	if err != nil {
		return nil, fmt.Errorf("failed to query credentials: %w", err)
	}
	defer rows.Close()

	var creds []*models.Credential
	for rows.Next() {
		var c models.Credential
		if err := rows.Scan(
			&c.ID,
			&c.VaultID,
			&c.FolderID,
			&c.Title,
			&c.URL,
			&c.EncryptedUsername,
			&c.EncryptedPassword,
			&c.EncryptedNotes,
			&c.CreatedAt,
			&c.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan credential: %w", err)
		}
		creds = append(creds, &c)
	}

	return creds, nil
}