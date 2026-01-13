/*
Royal Vault Backend API

File: main.go
Purpose: API server entry point
Tech: Golang, Fiber

This server handles authentication, vault storage,
and user management. It never handles plaintext secrets.
*/

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "Royal Vault API running",
		})
	})

	log.Fatal(app.Listen(":5000"))
}
