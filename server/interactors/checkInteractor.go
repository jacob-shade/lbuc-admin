package interactors

import (
	"fmt"

	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/model"
)

func CreateCheck(playerID uint, taskID uint) (model.Check, error) {
	// Get Player
	player, err := GetPlayer(playerID)
	if err != nil {
		return model.Check{}, fmt.Errorf("failed to get player: %w", err)
	}

	// Get Task
	task, err := GetTask(taskID)
	if err != nil {
		return model.Check{}, fmt.Errorf("failed to get task: %w", err)
	}

	// Create Check
	check, err := database.CreateCheck(player, task)
	if err != nil {
		return model.Check{}, fmt.Errorf("failed to create check: %w", err)
	}

	return check, nil
}

func GetCheck(playerId uint, taskId uint) (model.Check, error) {
	check, err := database.GetCheck(playerId, taskId)
	if err != nil {
		return model.Check{}, fmt.Errorf("failed to get check: %w", err)
	}
	return check, nil
}

func GetChecksForPlayer(playerID uint) ([]model.Check, error) {
	// Get Player
	player, err := GetPlayer(playerID)
	if err != nil {
		return []model.Check{}, fmt.Errorf("failed to get player: %w", err)
	}

	// Get Checks
	checks, err := database.GetChecksForPlayer(player)
	if err != nil {
		return []model.Check{}, fmt.Errorf("failed to get checks: %w", err)
	}

	return checks, nil
}

func GetChecksForTask(taskID uint) ([]model.Check, error) {
	// Get Task
	task, err := GetTask(taskID)
	if err != nil {
		return []model.Check{}, fmt.Errorf("failed to get task: %w", err)
	}

	// Get Checks
	checks, err := database.GetChecksForTask(task)
	if err != nil {
		return []model.Check{}, fmt.Errorf("failed to get checks: %w", err)
	}

	return checks, nil
}

func GetAllChecks() ([]model.Check, error) {
	checks, err := database.GetAllChecks()
	if err != nil {
		return []model.Check{}, fmt.Errorf("failed to get all checks: %w", err)
	}
	return checks, nil
}

func UpdateCheck(newCheck model.Check) error {
	// Get Check
	check, err := GetCheck(newCheck.PlayerID, newCheck.TaskID)
	if err != nil {
		return fmt.Errorf("failed to get check: %w", err)
	}

	// Update Check
	check.Checked = newCheck.Checked
	if err := database.UpdateCheck(check); err != nil {
		return fmt.Errorf("failed to update check: %w", err)
	}
	return nil
}

func DeleteCheck(check model.Check) error {
	// Get check from database
	check, err := GetCheck(check.PlayerID, check.TaskID)
	if err != nil {
		return fmt.Errorf("failed to get check: %w", err)
	}

	// Delete check from database
	if err := database.DeleteCheck(check); err != nil {
		return fmt.Errorf("failed to delete check: %w", err)
	}
	return nil
}
