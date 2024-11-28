package controllers

import (
	"github.com/jacobshade/lbuc-admin/server/initializers"
	"github.com/jacobshade/lbuc-admin/server/models"

	"github.com/gofiber/fiber/v2"
)

// Serialized Team
type Team struct {
	ID       uint     `json:"id"`
	TeamName string   `json:"team_name"`
	Players  []Player `json:"players"`
}

type TeamBasic struct {
	ID       uint   `json:"id"`
	TeamName string `json:"team_name"`
}

func CreateResponseTeam(teamModel models.Team) Team {
	players := make([]Player, len(teamModel.Players))
	for i, p := range teamModel.Players {
		players[i] = Player{
			ID: p.ID, PlayerName: p.PlayerName, NickName: p.NickName,
			Pronouns: p.Pronouns, Grade: p.Grade, Birthday: p.Birthday,
			PlayerEmail: p.PlayerEmail, ParentName: p.ParentName, ParentEmail: p.ParentEmail,
			ParentNumber: p.ParentNumber, Relationship: p.Relationship, Address: p.Address,
			MedicalNotes: p.MedicalNotes,
		}
	}
	return Team{ID: teamModel.ID, TeamName: teamModel.TeamName, Players: players}
}

// Create a new team with a given name and empty player list
func CreateTeam(c *fiber.Ctx) error {
	team := models.Team{}
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	initializers.DB.Create(&team)
	responseTeam := CreateResponseTeam(team)

	return c.Status(fiber.StatusOK).JSON(responseTeam)
}

// Gets all teams with ID and TeamName
func GetAllTeams(c *fiber.Ctx) error {
	teams := []models.Team{}
	if err := initializers.DB.Find(&teams).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	responseTeams := make([]TeamBasic, len(teams))
	for i, team := range teams {
		responseTeams[i] = TeamBasic{ID: team.ID, TeamName: team.TeamName}
	}

	return c.Status(fiber.StatusOK).JSON(responseTeams)
}

// Gets a team by ID with a list of all players on the team
func GetTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var team models.Team
	if err := initializers.DB.Preload("Players").First(&team, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Team not found"})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponseTeam(team))
}

func UpdateTeamName(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	team := models.Team{}
	if err := initializers.DB.First(&team, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Team not found"})
	}

	var updateData TeamBasic
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	team.TeamName = updateData.TeamName
	initializers.DB.Save(&team)

	return c.Status(fiber.StatusOK).JSON(TeamBasic{ID: team.ID, TeamName: team.TeamName})
}

func DeleteTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	team := models.Team{}
	if err := initializers.DB.First(&team, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Team not found"})
	}

	if err := initializers.DB.Delete(&team).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponseTeam(team))
}

func AddPlayerToTeam(c *fiber.Ctx) error {
	// Get team ID from params
	teamID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Parse request body
	type AddPlayersRequest struct {
		PlayerIDs []uint `json:"player_ids"`
	}
	var req AddPlayersRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Find team
	var team models.Team
	if err := initializers.DB.First(&team, teamID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Team not found"})
	}

	// Find all players
	var players []models.Player
	if err := initializers.DB.Find(&players, req.PlayerIDs).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "One or more players not found"})
	}

	// Add players to team
	if err := initializers.DB.Model(&team).Association("Players").Append(&players); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Get updated team with players
	if err := initializers.DB.Preload("Players").First(&team, teamID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponseTeam(team))
}

func RemovePlayerFromTeam(c *fiber.Ctx) error {
	// Get IDs from params
	teamID, err := c.ParamsInt("teamId") // Assuming route param is "teamId"
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid team ID"})
	}

	playerID, err := c.ParamsInt("playerId") // Assuming route param is "playerId"
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid player ID"})
	}

	// Find team
	var team models.Team
	if err := initializers.DB.First(&team, teamID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Team not found"})
	}

	// Find player
	var player models.Player
	if err := initializers.DB.First(&player, playerID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}

	// Remove player from team
	if err := initializers.DB.Model(&team).Association("Players").Delete(&player); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Get updated team with players
	if err := initializers.DB.Preload("Players").First(&team, teamID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponseTeam(team))
}
