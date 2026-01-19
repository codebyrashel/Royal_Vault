package services

import (
	"context"
	"fmt"

	"github.com/codebyrashel/Royal_Vault/server/internal/models"
	"github.com/codebyrashel/Royal_Vault/server/internal/repositories"
)

type SecurityQuestionService struct {
	vaultRepo             *repositories.VaultRepository
	credentialRepo        *repositories.CredentialRepository
	securityQuestionRepo  *repositories.SecurityQuestionRepository
}

func NewSecurityQuestionService(
	vaultRepo *repositories.VaultRepository,
	credentialRepo *repositories.CredentialRepository,
	securityQuestionRepo *repositories.SecurityQuestionRepository,
) *SecurityQuestionService {
	return &SecurityQuestionService{
		vaultRepo:            vaultRepo,
		credentialRepo:       credentialRepo,
		securityQuestionRepo: securityQuestionRepo,
	}
}

// Ensure credential belongs to user's vault
func (s *SecurityQuestionService) validateCredentialOwnership(
	ctx context.Context,
	userID int64,
	credentialID int64,
) error {
	vault, err := s.vaultRepo.GetVaultByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to load vault: %w", err)
	}
	if vault == nil {
		return fmt.Errorf("vault not found for user")
	}

	creds, err := s.credentialRepo.ListCredentialsByVaultID(ctx, vault.ID)
	if err != nil {
		return fmt.Errorf("failed to load credentials for vault: %w", err)
	}

	found := false
	for _, c := range creds {
		if c.ID == credentialID {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("credential does not belong to user")
	}

	return nil
}

func (s *SecurityQuestionService) Create(
	ctx context.Context,
	userID int64,
	credentialID int64,
	encryptedQuestion string,
	encryptedAnswer string,
) (*models.SecurityQuestion, error) {
	if err := s.validateCredentialOwnership(ctx, userID, credentialID); err != nil {
		return nil, err
	}

	return s.securityQuestionRepo.CreateSecurityQuestion(
		ctx,
		credentialID,
		encryptedQuestion,
		encryptedAnswer,
	)
}

func (s *SecurityQuestionService) ListForCredential(
	ctx context.Context,
	userID int64,
	credentialID int64,
) ([]*models.SecurityQuestion, error) {
	if err := s.validateCredentialOwnership(ctx, userID, credentialID); err != nil {
		return nil, err
	}

	return s.securityQuestionRepo.ListByCredentialID(ctx, credentialID)
}

func (s *SecurityQuestionService) Delete(
	ctx context.Context,
	userID int64,
	id int64,
	credentialID int64,
) error {
	if err := s.validateCredentialOwnership(ctx, userID, credentialID); err != nil {
		return err
	}

	return s.securityQuestionRepo.DeleteByID(ctx, id)
}