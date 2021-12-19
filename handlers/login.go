package handlers

import (
	"auth/entities"
	"auth/lib"
	"auth/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	usernameKey = "username"
	passwordKey = "password"
)

func Login(collection *mongo.Collection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.Locals(usernameKey)
		password := c.Locals(passwordKey).(string)

		// Search if user has already registered
		query := bson.D{{Key: "username", Value: username}}
		existingUser := collection.FindOne(c.Context(), query)

		if existingUser.Err() == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": fiber.StatusNotFound,
				"error": &fiber.Map{
					"message": lib.UserNotFound,
				},
			})
		}

		userFromDB := &entities.User{}
		existingUser.Decode(userFromDB)

		compareError := utils.ComparePassword([]byte(userFromDB.Password), []byte(password))

		if compareError != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  lib.InvalidCredentials,
			})
		}

		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"data":   userFromDB,
		})
	}
}
