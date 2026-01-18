package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codebyrashel/Royal_Vault/server/internal/models"
)

type FolderRepository struct {
	db *sql.DB
}

func NewFolderRepository(db *sql.DB) *FolderRepository {
	return &FolderRepository{db: db}
}

func (r *FolderRepository) CreateFolder(
	ctx context.Context,
	userID int64,
	name string,
) (*models.Folder, error) {
	query := `
		INSERT INTO folders (user_id, name, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, user_id, name, created_at, updated_at
	`

	row := r.db.QueryRowContext(ctx, query, userID, name)

	var folder models.Folder
	if err := row.Scan(
		&folder.ID,
		&folder.UserID,
		&folder.Name,
		&folder.CreatedAt,
		&folder.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to insert folder: %w", err)
	}

	return &folder, nil
}

func (r *FolderRepository) ListFoldersByUserID(
	ctx context.Context,
	userID int64,
) ([]*models.Folder, error) {
	query := `
		SELECT id, user_id, name, created_at, updated_at
		FROM folders
		WHERE user_id = $1
		ORDER BY name ASC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query folders: %w", err)
	}
	defer rows.Close()

	var folders []*models.Folder
	for rows.Next() {
		var f models.Folder
		if err := rows.Scan(
			&f.ID,
			&f.UserID,
			&f.Name,
			&f.CreatedAt,
			&f.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan folder: %w", err)
		}
		folders = append(folders, &f)
	}

	return folders, nil
}