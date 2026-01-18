package services

import (
	"context"
	"fmt"

	"github.com/codebyrashel/Royal_Vault/server/internal/models"
	"github.com/codebyrashel/Royal_Vault/server/internal/repositories"
)

type CredentialService struct {
	vaultRepo      *repositories.VaultRepository
	credentialRepo *repositories.CredentialRepository
}

func NewCredentialService(
	vaultRepo *repositories.VaultRepository,
	credentialRepo *repositories.CredentialRepository,
) *CredentialService {
	return &CredentialService{
		vaultRepo:      vaultRepo,
		credentialRepo: credentialRepo,
	}
}

func (s *CredentialService) CreateCredential(
	ctx context.Context,
	userID int64,
	folderID *int64,
	title string,
	url *string,
	encryptedUsername string,
	encryptedPassword string,
	encryptedNotes *string,
) (*models.Credential, error) {
	vault, err := s.vaultRepo.GetVaultByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to load vault: %w", err)
	}
	if vault == nil {
		return nil, fmt.Errorf("vault not found for user")
	}

	return s.credentialRepo.CreateCredential(
		ctx,
		vault.ID,
		folderID,
		title,
		url,
		encryptedUsername,
		encryptedPassword,
		encryptedNotes,
	)
}

func (s *CredentialService) ListCredentials(
	ctx context.Context,
	userID int64,
) ([]*models.Credential, error) {
	vault, err := s.vaultRepo.GetVaultByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to load vault: %w", err)
	}
	if vault == nil {
		return nil, fmt.Errorf("vault not found for user")
	}

	return s.credentialRepo.ListCredentialsByVaultID(ctx, vault.ID)
}