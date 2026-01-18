package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/codebyrashel/Royal_Vault/server/internal/models"
	"github.com/codebyrashel/Royal_Vault/server/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *repositories.UserRepository
	vaultRepo *repositories.VaultRepository
	jwtKey    []byte
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	vaultRepo *repositories.VaultRepository,
) *AuthService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "change-me-in-production"
	}

	return &AuthService{
		userRepo:  userRepo,
		vaultRepo: vaultRepo,
		jwtKey:    []byte(secret),
	}
}

func (s *AuthService) Signup(
	ctx context.Context,
	email, password, encryptedVaultKey, salt string,
) (*models.User, error) {
	existing, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("user with this email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := s.userRepo.CreateUser(ctx, email, string(hash))
	if err != nil {
		return nil, err
	}

	_, err = s.vaultRepo.CreateVault(ctx, user.ID, encryptedVaultKey, salt)
	if err != nil {
		return nil, fmt.Errorf("failed to create vault: %w", err)
	}

	return user, nil
}

type LoginResult struct {
	User             *models.User
	Token            string
	EncryptedVaultKey string
	Salt             string
}

func (s *AuthService) Login(
	ctx context.Context,
	email, password string,
) (*LoginResult, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	vault, err := s.vaultRepo.GetVaultByUserID(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to load vault: %w", err)
	}
	if vault == nil {
		return nil, fmt.Errorf("vault not found for user")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &LoginResult{
		User:             user,
		Token:            token,
		EncryptedVaultKey: vault.EncryptedVaultKey,
		Salt:             vault.Salt,
	}, nil
}

func (s *AuthService) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"userId": user.ID,
		"email":  user.Email,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(s.jwtKey)
}