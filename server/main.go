package main

import (
	"log"

	"github.com/jacobshade/lbuc-admin/server/initializers"
	"github.com/jacobshade/lbuc-admin/server/routes"

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
