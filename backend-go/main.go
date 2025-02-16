package main

import (
	"dynamic-pricing-api/backend-go/database"
	"dynamic-pricing-api/backend-go/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Connect to local database, and migrate current schemas
	database.ConnectDB()
	database.MigrateDB()

	log.Println("Dynamic Pricing API is running...")

	// Create the application instance
	app := fiber.New()

	// Use logging middleware to log every request
	app.Use(logger.New())

	// Set up endpoint routes
	routes.SetupRoutes(app)

	// Start the server
	app.Listen(":8080")
}
