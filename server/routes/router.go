package routes

import (
	"github.com/jacobshade/lbuc-admin/server/handler"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	// Healthcheck endpoint
	api.Get("/", handler.HealthCheck)

	// Auth endpoints /api/auth
	auth := api.Group("/auth")
	auth.Get("/", handler.Auth)
	auth.Get("/google/callback", handler.Callback)

	// User endpoints /api/user
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)
	user.Get("/", handler.GetUsers)
	user.Get("/:id", handler.GetUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	// Player endpoints /api/player
	player := api.Group("/player")
	player.Post("/", handler.CreatePlayer)
	player.Get("/", handler.GetAllPlayers)
	player.Get("/:id", handler.GetPlayer)
	player.Put("/:id", handler.UpdatePlayer)
	player.Delete("/:id", handler.DeletePlayer)

	// Team endpoints /api/team
	team := api.Group("/team")
	team.Post("/", handler.CreateTeam)
	team.Get("/", handler.GetAllTeams)
	team.Get("/:id", handler.GetTeam)
	team.Put("/:id", handler.UpdateTeamName)
	team.Delete("/:id", handler.DeleteTeam)
	team.Post("/:id/player", handler.AddPlayerToTeam)
	team.Delete("/:teamId/player/:playerId", handler.RemovePlayerFromTeam)
}
