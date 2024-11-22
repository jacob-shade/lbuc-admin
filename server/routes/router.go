package routes

import (
	"lbuc-admin/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Healthcheck endpoint
	app.Get("/api", controllers.HealthCheck)

	// User endpoints
	app.Post("/api/user", controllers.CreateUser)
	app.Get("/api/user", controllers.GetUsers)
	app.Get("/api/user/:id", controllers.GetUser)
	app.Put("/api/user/:id", controllers.UpdateUser)
	app.Delete("/api/user/:id", controllers.DeleteUser)
}
