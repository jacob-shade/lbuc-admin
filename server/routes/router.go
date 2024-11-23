package routes

import (
	"github.com/jacobshade/lbuc-admin/server/controllers"

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

	// Player endpoints
	app.Post("/api/player", controllers.CreatePlayer)
	app.Get("/api/player", controllers.GetAllPlayers)
	app.Get("/api/player/:id", controllers.GetPlayer)
	app.Put("/api/player/:id", controllers.UpdatePlayer)
	app.Delete("/api/player/:id", controllers.DeletePlayer)

	// Team endpoints
	app.Post("/api/team", controllers.CreateTeam)
	app.Get("/api/team", controllers.GetAllTeams)
	app.Get("/api/team/:id", controllers.GetTeam)
	app.Put("/api/team/:id", controllers.UpdateTeamName)
	app.Delete("/api/team/:id", controllers.DeleteTeam)
	app.Post("/api/team/:id/player", controllers.AddPlayerToTeam)
	app.Delete("/api/team/:teamId/player/:playerId", controllers.RemovePlayerFromTeam)
}
