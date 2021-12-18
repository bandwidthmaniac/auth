package handlers

import (
	"auth/entities"
	"auth/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(collection *mongo.Collection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body *entities.User
		bodyParserError := c.BodyParser(&body)

		// Payload is unprocessable, respond with a bad response
		if bodyParserError != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"error": bodyParserError,
			})
		}

		// Search if user has already registered
		query := bson.D{{Key: "username", Value: body.Username}}
		existingUser := collection.FindOne(c.Context(), query)

		userFromDb := body
		existingUser.Decode(userFromDb)

		// If user with provided username already exists,
		// respond with bad request
		if existingUser.Err() != mongo.ErrNoDocuments {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": fiber.StatusBadRequest,
				"error": fiber.Map{
					"message": "User is already registered",
				},
			})
		}

		hashedPassword, hashingError := utils.HashPassword([]byte(body.Password))

		if hashingError != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": hashingError,
			})
		}

		user := &entities.User{}
		user.Password = hashedPassword

		insertionResult, err := collection.InsertOne(c.Context(), user)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": fiber.StatusInternalServerError,
				"error":  err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Registration successful",
			"data": fiber.Map{
				"id": insertionResult.InsertedID,
			},
		})
	}
}
