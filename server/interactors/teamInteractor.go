package interactors

import (
	"fmt"

	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/model"
)

func CreateTeam(team model.Team) (model.Team, error) {
	team, err := database.CreateTeam(team)
	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func GetTeam(id uint) (model.Team, error) {
	team, err := database.GetTeam(id)
	if err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func GetAllTeams() ([]model.Team, error) {
	teams, err := database.GetAllTeams()
	if err != nil {
		return []model.Team{}, err
	}
	return teams, nil
}

func UpdateTeamName(id uint, teamName string) error {
	team, err := GetTeam(id)
	if err != nil {
		return err
	}

	team.TeamName = teamName
	if err := database.UpdateTeam(team); err != nil {
		return err
	}
	return nil
}

func DeleteTeam(id uint) error {
	// Find team
	team, err := GetTeam(id)
	if err != nil {
		return fmt.Errorf("failed to find team: %w", err)
	}

	// Delete team
	if err := database.DeleteTeam(team); err != nil {
		return fmt.Errorf("failed to delete team: %w", err)
	}
	return nil
}

func AddPlayersToTeam(teamID uint, playerIDs []uint) error {
	// Find team
	team, err := GetTeam(teamID)
	if err != nil {
		return fmt.Errorf("failed to find team: %w", err)
	}

	// Find all players
	players, err := GetListOfPlayers(playerIDs)
	if err != nil {
		return fmt.Errorf("failed to find players: %w", err)
	}

	// Add players to team
	if err := database.AddPlayersToTeam(team, players); err != nil {
		return fmt.Errorf("failed to add players to team: %w", err)
	}

	return nil
}

func RemovePlayerFromTeam(teamID uint, playerID uint) error {
	// Find team
	team, err := GetTeam(teamID)
	if err != nil {
		return fmt.Errorf("failed to find team: %w", err)
	}

	// Find player
	player, err := GetPlayer(playerID)
	if err != nil {
		return fmt.Errorf("failed to find player: %w", err)
	}

	// Remove player from team
	if err := database.RemovePlayerFromTeam(team, player); err != nil {
		return fmt.Errorf("failed to remove player from team: %w", err)
	}

	return nil
}

func CreateTaskForTeam(teamID uint, description string) (model.Team, error) {
	// Find team
	team, err := GetTeam(teamID)
	if err != nil {
		return model.Team{}, fmt.Errorf("failed to find team: %w", err)
	}

	// Create task
	task, err := CreateTask(description, team)
	if err != nil {
		return model.Team{}, fmt.Errorf("failed to create task: %w", err)
	}

	// Get all players on the team
	players, err := database.GetAllPlayersOnTeam(teamID)
	if err != nil {
		return model.Team{}, fmt.Errorf("failed to find players: %w", err)
	}

	// Add check for each player on the team
	for _, player := range players {
		CreateCheck(player.ID, task.ID)
	}
	return team, nil
}

func RemoveTaskFromTeam(teamID uint, taskID uint) error {
	// Find team
	team, err := GetTeam(teamID)
	if err != nil {
		return fmt.Errorf("failed to find team: %w", err)
	}

	// Find task
	task, err := GetTask(taskID)
	if err != nil {
		return fmt.Errorf("failed to find task: %w", err)
	}

	// Remove task from team
	if err := database.RemoveTaskFromTeam(team, task); err != nil {
		return fmt.Errorf("failed to remove task from team: %w", err)
	}

	return nil
}
