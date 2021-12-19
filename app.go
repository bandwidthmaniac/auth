package main

import (
	"auth/config"
	"auth/entities"
	"auth/handlers"
	"auth/middleware"

	"context"
	"flag"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mg       entities.MongoInstance
	dbName   = config.Get("DB_NAME")
	mongoURI = config.Get("MONGO_URI") + dbName
	prod     = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connect with MongoDB
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	userCollection := mg.Db.Collection(config.Get("DB_COLLECTION"))

	// Initialize a new Fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod,
	})

	// Middleware
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(etag.New())
	app.Use(cors.New())
	app.Use(requestid.New())

	// Group /v1 endpoint.
	v1 := app.Group("/v1")

	v1.Post("/register", handlers.Register(userCollection))
	v1.Put("/login", middleware.ValidateAuthPayload, handlers.Login(userCollection))

	// Handle other 404 routes
	app.Use(handlers.NotFound)

	// Configure port to listen on
	port := config.Get("PORT")

	if config.Get("PORT") == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
func Connect() error {
	client, clientError := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if clientError != nil {
		return clientError
	}

	connectionError := client.Connect(ctx)
	db := client.Database(dbName)

	if connectionError != nil {
		return clientError
	}

	mg = entities.MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}
