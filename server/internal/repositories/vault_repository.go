package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codebyrashel/Royal_Vault/server/internal/models"
)

type VaultRepository struct {
	db *sql.DB
}

func NewVaultRepository(db *sql.DB) *VaultRepository {
	return &VaultRepository{db: db}
}

func (r *VaultRepository) CreateVault(
	ctx context.Context,
	userID int64,
	encryptedVaultKey string,
	salt string,
) (*models.Vault, error) {
	query := `
		INSERT INTO vaults (user_id, encrypted_vault_key, salt, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, user_id, encrypted_vault_key, salt, created_at, updated_at
	`

	row := r.db.QueryRowContext(ctx, query, userID, encryptedVaultKey, salt)

	var vault models.Vault
	if err := row.Scan(
		&vault.ID,
		&vault.UserID,
		&vault.EncryptedVaultKey,
		&vault.Salt,
		&vault.CreatedAt,
		&vault.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to insert vault: %w", err)
	}

	return &vault, nil
}

func (r *VaultRepository) GetVaultByUserID(
	ctx context.Context,
	userID int64,
) (*models.Vault, error) {
	query := `
		SELECT id, user_id, encrypted_vault_key, salt, created_at, updated_at
		FROM vaults
		WHERE user_id = $1
	`

	row := r.db.QueryRowContext(ctx, query, userID)

	var vault models.Vault
	if err := row.Scan(
		&vault.ID,
		&vault.UserID,
		&vault.EncryptedVaultKey,
		&vault.Salt,
		&vault.CreatedAt,
		&vault.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query vault by user id: %w", err)
	}

	return &vault, nil
}