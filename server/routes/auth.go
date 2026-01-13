/*
Auth Routes

Defines authentication endpoints.
*/

package routes

import (
	"royal-vault/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}

