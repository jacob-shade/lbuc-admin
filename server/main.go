package main

import (
	"log"

	"github.com/jacobshade/lbuc-admin/server/config"
	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.LoadEnvVariables()
	database.ConnectToDatabase()
	database.SetupSessionStore()
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(":5000"))
	// go func() {
	// 	log.Fatal(app.Listen(":5000"))
	// }()
}
