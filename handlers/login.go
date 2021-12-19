package handlers

import (
	"auth/config"
	"auth/entities"
	"auth/lib"
	"auth/utils"

	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/golang-jwt/jwt"
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

		if !utils.ComparePassword([]byte(userFromDB.Password), []byte(password)) {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": fiber.StatusBadRequest,
				"error":  lib.InvalidCredentials,
			})
		}

		jwtSecret := config.Get("JWT_SECRET")
		encodedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  userFromDB.ID,
			"iat": time.Now(),
		})

		signedToken, errTokenSign := encodedToken.SignedString([]byte(jwtSecret))

		if errTokenSign != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": fiber.StatusInternalServerError,
				"error":  fiber.ErrInternalServerError,
			})
		}

		// User data will be sent as response
		userData := map[string]string{
			"id":    userFromDB.ID.Hex(),
			"token": signedToken,
		}

		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"data":   userData,
		})
	}
}
