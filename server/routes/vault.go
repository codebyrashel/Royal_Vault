package routes

import (
	"royal-vault/middleware"

	"github.com/gofiber/fiber/v2"
)

func VaultRoutes(app *fiber.App) {
	vault := app.Group("/api/vault", middleware.Protected())

	vault.Get("/me", func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(fiber.Map{
			"user": user,
		})
	})
}
