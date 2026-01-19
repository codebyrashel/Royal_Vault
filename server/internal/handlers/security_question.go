package handlers

import (
	"net/http"
	"strconv"

	"github.com/codebyrashel/Royal_Vault/server/internal/middleware"
	"github.com/codebyrashel/Royal_Vault/server/internal/services"
	"github.com/gin-gonic/gin"
)

type SecurityQuestionHandler struct {
	securityQuestionService *services.SecurityQuestionService
}

func NewSecurityQuestionHandler(
	securityQuestionService *services.SecurityQuestionService,
) *SecurityQuestionHandler {
	return &SecurityQuestionHandler{securityQuestionService: securityQuestionService}
}

type createSecurityQuestionRequest struct {
	CredentialID      int64  `json:"credentialId" binding:"required"`
	EncryptedQuestion string `json:"encryptedQuestion" binding:"required"`
	EncryptedAnswer   string `json:"encryptedAnswer" binding:"required"`
}

func (h *SecurityQuestionHandler) Create(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	var req createSecurityQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	sq, err := h.securityQuestionService.Create(
		c.Request.Context(),
		userID,
		req.CredentialID,
		req.EncryptedQuestion,
		req.EncryptedAnswer,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                sq.ID,
		"credentialId":      sq.CredentialID,
		"encryptedQuestion": sq.EncryptedQuestion,
		"encryptedAnswer":   sq.EncryptedAnswer,
	})
}

func (h *SecurityQuestionHandler) ListForCredential(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	credentialIDParam := c.Param("credentialId")
	credentialID, err := strconv.ParseInt(credentialIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credential ID",
		})
		return
	}

	sqs, err := h.securityQuestionService.ListForCredential(
		c.Request.Context(),
		userID,
		credentialID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var resp []gin.H
	for _, sq := range sqs {
		resp = append(resp, gin.H{
			"id":                sq.ID,
			"credentialId":      sq.CredentialID,
			"encryptedQuestion": sq.EncryptedQuestion,
			"encryptedAnswer":   sq.EncryptedAnswer,
		})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *SecurityQuestionHandler) Delete(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	credentialIDParam := c.Query("credentialId")
	credentialID, err := strconv.ParseInt(credentialIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credential ID",
		})
		return
	}

	if err := h.securityQuestionService.Delete(
		c.Request.Context(),
		userID,
		id,
		credentialID,
	); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}