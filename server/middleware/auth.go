/*
JWT Authentication Middleware

Verifies JWT token and authorizes requests.
*/

package middleware

import (
	"royal-vault/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: config.JWTSecret,
	})
}
