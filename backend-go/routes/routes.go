package routes

import (
	"dynamic-pricing-api/backend-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/products", controllers.GetProducts)
	app.Post("/products", controllers.CreateProduct)
	app.Put("/products/:id", controllers.UpdateProductPrice)
}
