package controllers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"

	"royal-vault/config"
	"royal-vault/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func generateSalt() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Register(c *fiber.Ctx) error {
	var input RegisterInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	salt, err := generateSalt()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate salt"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password+salt), 12)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	var user models.User

	err = config.DB.QueryRow(
		"INSERT INTO users(email, password_hash, salt) VALUES($1,$2,$3) RETURNING id,email,password_hash,salt",
		input.Email,
		string(hash),
		salt,
	).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Salt)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(500).JSON(fiber.Map{"error": "Insert failed"})
		}
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user.Email,
	})
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Login endpoint coming soon",
	})
}
