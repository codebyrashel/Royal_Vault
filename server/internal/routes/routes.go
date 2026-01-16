package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/codebyrashel/Royal_Vault/server/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	healthHandler := handlers.NewHealthHandler()

	// Health check route
	router.GET("/health", healthHandler.GetHealth)

	// Future: auth, vault, credentials, security questions, etc.

	return router
}