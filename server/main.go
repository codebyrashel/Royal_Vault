package main

import (
	"log"

	"royal-vault/config"
	"royal-vault/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()

	routes.AuthRoutes(app)
	routes.VaultRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "Royal Vault API running"})
	})

	log.Fatal(app.Listen(":5000"))
}
