package middleware

import (
	"auth/entities"
	"auth/lib"

	"github.com/gofiber/fiber/v2"
)

// Validates the auth payload such as `username` & `password`
// by enforcing constraints.
func ValidateAuthPayload(c *fiber.Ctx) error {
	var expectedPayload *entities.User

	if parserError := c.BodyParser(&expectedPayload); parserError != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": fiber.StatusUnprocessableEntity,
			"error":  parserError,
		})
	}

	// Length constraints
	if len(expectedPayload.Username) < 6 || len(expectedPayload.Password) < 6 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"status": fiber.StatusUnprocessableEntity,
			"error":  fiber.NewError(fiber.StatusUnprocessableEntity, lib.PayloadLengthError),
		})
	}

	c.Locals("username", expectedPayload.Username)
	c.Locals("password", expectedPayload.Password)
	return c.Next()
}
