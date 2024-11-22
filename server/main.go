package main

import (
	"lbuc-admin/initializers"
	"lbuc-admin/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen(":5000"))
}
