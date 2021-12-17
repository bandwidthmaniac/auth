package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	// password := []byte("Pa$$w0rD$ecr3t")
	// hashedPassowrd, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	return c.JSON(fiber.Map{
		"status": "Registration successful.",
	})
}
