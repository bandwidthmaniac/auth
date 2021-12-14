package handlers

import (
	"auth/entities"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	body := new(entities.User)

	err := c.BodyParser(body)

	if err != nil {
		return c.JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"status": "Registration successful.",
	})
}
