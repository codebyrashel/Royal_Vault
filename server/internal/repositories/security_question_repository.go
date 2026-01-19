package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/codebyrashel/Royal_Vault/server/internal/models"
)

type SecurityQuestionRepository struct {
	db *sql.DB
}

func NewSecurityQuestionRepository(db *sql.DB) *SecurityQuestionRepository {
	return &SecurityQuestionRepository{db: db}
}

func (r *SecurityQuestionRepository) CreateSecurityQuestion(
	ctx context.Context,
	credentialID int64,
	encryptedQuestion string,
	encryptedAnswer string,
) (*models.SecurityQuestion, error) {
	query := `
		INSERT INTO security_questions (
			credential_id,
			encrypted_question,
			encrypted_answer,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id, credential_id, encrypted_question, encrypted_answer, created_at, updated_at
	`

	row := r.db.QueryRowContext(ctx, query, credentialID, encryptedQuestion, encryptedAnswer)

	var sq models.SecurityQuestion
	if err := row.Scan(
		&sq.ID,
		&sq.CredentialID,
		&sq.EncryptedQuestion,
		&sq.EncryptedAnswer,
		&sq.CreatedAt,
		&sq.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to insert security question: %w", err)
	}

	return &sq, nil
}

func (r *SecurityQuestionRepository) ListByCredentialID(
	ctx context.Context,
	credentialID int64,
) ([]*models.SecurityQuestion, error) {
	query := `
		SELECT id, credential_id, encrypted_question, encrypted_answer, created_at, updated_at
		FROM security_questions
		WHERE credential_id = $1
		ORDER BY id ASC
	`

	rows, err := r.db.QueryContext(ctx, query, credentialID)
	if err != nil {
		return nil, fmt.Errorf("failed to query security questions: %w", err)
	}
	defer rows.Close()

	var result []*models.SecurityQuestion
	for rows.Next() {
		var sq models.SecurityQuestion
		if err := rows.Scan(
			&sq.ID,
			&sq.CredentialID,
			&sq.EncryptedQuestion,
			&sq.EncryptedAnswer,
			&sq.CreatedAt,
			&sq.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan security question: %w", err)
		}
		result = append(result, &sq)
	}

	return result, nil
}

func (r *SecurityQuestionRepository) DeleteByID(
	ctx context.Context,
	id int64,
) error {
	query := `DELETE FROM security_questions WHERE id = $1`
	if _, err := r.db.ExecContext(ctx, query, id); err != nil {
		return fmt.Errorf("failed to delete security question: %w", err)
	}
	return nil
}