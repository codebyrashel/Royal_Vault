package routes

import (
	"database/sql"

	"github.com/codebyrashel/Royal_Vault/server/internal/handlers"
	"github.com/codebyrashel/Royal_Vault/server/internal/repositories"
	"github.com/codebyrashel/Royal_Vault/server/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	healthHandler := handlers.NewHealthHandler()

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := NewAuthHandlerWrapper(authService)

	// Health check
	router.GET("/health", healthHandler.GetHealth)

	// Auth
	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
	}

	return router
}

// helper to keep handler construction in one place
func NewAuthHandlerWrapper(authService *services.AuthService) *handlers.AuthHandler {
	return handlers.NewAuthHandler(authService)
}