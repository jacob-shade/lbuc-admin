package handler

import (
	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/interactors"
	"github.com/jacobshade/lbuc-admin/server/model"

	"github.com/gofiber/fiber/v2"
)

// Serialized Team
type Team struct {
	ID       uint     `json:"id"`
	TeamName string   `json:"team_name"`
	Players  []Player `json:"players"`
	Tasks    []Task   `json:"tasks"`
}

type TeamBasic struct {
	ID       uint   `json:"id"`
	TeamName string `json:"team_name"`
}

func CreateResponseTeam(teamModel model.Team) Team {
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
	tasks := make([]Task, len(teamModel.Tasks))
	for i, t := range teamModel.Tasks {
		tasks[i] = Task{ID: t.ID, Description: t.Description}
	}
	return Team{ID: teamModel.ID, TeamName: teamModel.TeamName, Players: players, Tasks: tasks}
}

func CreateTeam(c *fiber.Ctx) error {
	team := model.Team{}
	if err := c.BodyParser(&team); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	team, err := interactors.CreateTeam(team)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	responseTeam := CreateResponseTeam(team)
	return c.Status(fiber.StatusOK).JSON(responseTeam)
}

func GetAllTeams(c *fiber.Ctx) error {
	teams, err := interactors.GetAllTeams()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	responseTeams := make([]TeamBasic, len(teams))
	for i, team := range teams {
		responseTeams[i] = TeamBasic{ID: team.ID, TeamName: team.TeamName}
	}

	return c.Status(fiber.StatusOK).JSON(responseTeams)
}

func GetTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var team model.Team
	if err := database.DB.Preload("Players").Preload("Tasks").First(&team, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Team not found"})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponseTeam(team))
}

func UpdateTeamName(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updateData TeamBasic
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := interactors.UpdateTeamName(uint(id), updateData.TeamName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Team name updated successfully"})
}

// does this handle deleting a players assiciation with a team? tasks
func DeleteTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := interactors.DeleteTeam(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Team deleted successfully"})
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

	if err := interactors.AddPlayersToTeam(uint(teamID), req.PlayerIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Players added to team successfully"})
}

func RemovePlayerFromTeam(c *fiber.Ctx) error {
	// Get IDs from params
	teamID, err := c.ParamsInt("teamId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid team ID"})
	}

	playerID, err := c.ParamsInt("playerId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid player ID"})
	}

	if err := interactors.RemovePlayerFromTeam(uint(teamID), uint(playerID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Player removed from team successfully"})
}

func CreateTaskForTeam(c *fiber.Ctx) error {
	// Get team ID
	teamID, err := c.ParamsInt("teamId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Get task
	var task model.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	team, err := interactors.CreateTaskForTeam(uint(teamID), task.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponseTeam(team))
}

// TODO: no DATABASE USAGE
func RemoveTaskFromTeam(c *fiber.Ctx) error {
	// Get team ID
	teamID, err := c.ParamsInt("teamId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Find team, make sure it exists
	var team model.Team
	if err := database.DB.First(&team, teamID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Team not found"})
	}

	// Get task id
	var task model.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Remove task from team
	database.DB.Model(&team).Association("Tasks").Delete(&task)

	// TODO: Fix team response
	return c.Status(fiber.StatusOK).JSON(CreateResponseTeam(team))
}
