package handlers

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)

	return c.JSON(fiber.Map{
		"error":   404,
		"message": "Not found",
	})
}
