package handlers

import (
	"net/http"

	"github.com/codebyrashel/Royal_Vault/server/internal/middleware"
	"github.com/codebyrashel/Royal_Vault/server/internal/services"
	"github.com/gin-gonic/gin"
)

type CredentialHandler struct {
	credentialService *services.CredentialService
}

func NewCredentialHandler(credentialService *services.CredentialService) *CredentialHandler {
	return &CredentialHandler{credentialService: credentialService}
}

type createCredentialRequest struct {
	FolderID          *int64  `json:"folderId"`
	Title             string  `json:"title" binding:"required"`
	URL               *string `json:"url"`
	EncryptedUsername string  `json:"encryptedUsername" binding:"required"`
	EncryptedPassword string  `json:"encryptedPassword" binding:"required"`
	EncryptedNotes    *string `json:"encryptedNotes"`
}

func (h *CredentialHandler) CreateCredential(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	var req createCredentialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	cred, err := h.credentialService.CreateCredential(
		c.Request.Context(),
		userID,
		req.FolderID,
		req.Title,
		req.URL,
		req.EncryptedUsername,
		req.EncryptedPassword,
		req.EncryptedNotes,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                cred.ID,
		"title":             cred.Title,
		"url":               cred.URL,
		"encryptedUsername": cred.EncryptedUsername,
		"encryptedPassword": cred.EncryptedPassword,
		"encryptedNotes":    cred.EncryptedNotes,
	})
}

func (h *CredentialHandler) ListCredentials(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	creds, err := h.credentialService.ListCredentials(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var resp []gin.H
	for _, cred := range creds {
		resp = append(resp, gin.H{
			"id":                cred.ID,
			"title":             cred.Title,
			"url":               cred.URL,
			"encryptedUsername": cred.EncryptedUsername,
			"encryptedPassword": cred.EncryptedPassword,
			"encryptedNotes":    cred.EncryptedNotes,
			"folderId":          cred.FolderID,
		})
	}

	c.JSON(http.StatusOK, resp)
}