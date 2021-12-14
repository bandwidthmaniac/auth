package main

import (
	"auth/database"
	"auth/handlers"

	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/joho/godotenv"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.Connect()

	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(cors.New())

	// Group /v1 endpoint.
	v1 := app.Group("/v1")

	v1.Get("/register", handlers.Register)
	v1.Get("/login", handlers.Login)

	// Handle other 404 routes
	app.Use(handlers.NotFound)

	log.Fatal(app.Listen(*port))
}
