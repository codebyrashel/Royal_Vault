package routes

import (
	"database/sql"
	"time"

	"github.com/codebyrashel/Royal_Vault/server/internal/handlers"
	"github.com/codebyrashel/Royal_Vault/server/internal/repositories"
	"github.com/codebyrashel/Royal_Vault/server/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	healthHandler := handlers.NewHealthHandler()

	userRepo := repositories.NewUserRepository(db)
	vaultRepo := repositories.NewVaultRepository(db)
	authService := services.NewAuthService(userRepo, vaultRepo)
	authHandler := handlers.NewAuthHandler(authService)

	router.GET("/health", healthHandler.GetHealth)

	auth := router.Group("/auth")
	{
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/login", authHandler.Login)
	}

	return router
}