package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobshade/lbuc-admin/server/handler"
	"github.com/jacobshade/lbuc-admin/server/middleware"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	// Healthcheck endpoint
	api.Get("/", handler.CheckStatus)

	// Auth endpoints /api/auth (public endpoints)
	auth := api.Group("/auth")
	auth.Get("/", handler.GoogleLogin)
	auth.Get("/google/callback", handler.GoogleCallback)
	auth.Get("/session", handler.GetSession)
	auth.Post("/signout", handler.Signout)

	// Protected endpoints
	api.Use(middleware.AuthRequired())

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
	player.Get("/:id/checks", handler.GetChecksForPlayer)

	// Team endpoints /api/team
	team := api.Group("/team")
	team.Post("/", handler.CreateTeam)
	team.Get("/", handler.GetAllTeams)
	team.Get("/:id", handler.GetTeam)
	team.Put("/:id", handler.UpdateTeamName)
	team.Delete("/:id", handler.DeleteTeam)
	team.Post("/:id/player", handler.AddPlayerToTeam)
	team.Delete("/:teamId/player/:playerId", handler.RemovePlayerFromTeam)
	team.Post("/:teamId/task", handler.AddTaskToTeam)
	team.Delete("/:teamId/task", handler.RemoveTaskFromTeam)

	// Check endpoints
	api.Get("/task/:id/checks", handler.GetChecksForTask)
	api.Post("/check", handler.UpdateCheck)
}
