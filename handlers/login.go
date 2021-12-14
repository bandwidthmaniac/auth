package handlers

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"user": fiber.Map{
			"name": "User-1",
		},
	})
}
