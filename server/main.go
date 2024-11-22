package main

import (
	"lbuc-admin/controllers"
	"lbuc-admin/initializers"
	"log"

	"github.com/gofiber/fiber/v3"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/api", controllers.Welcome)

	log.Fatal(app.Listen(":5000"))
}
