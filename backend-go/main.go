package main

import (
	"dynamic-pricing-api/backend-go/database"
	"dynamic-pricing-api/backend-go/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	database.MigrateDB()

	log.Println("Dynamic Pricing API is running...")

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
