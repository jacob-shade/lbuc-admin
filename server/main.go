package main

import (
	"log"

	"github.com/jacobshade/lbuc-admin/server/initializers"
	"github.com/jacobshade/lbuc-admin/server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)

	log.Fatal(app.Listen(":5000"))
}
