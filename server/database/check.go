package database

import "github.com/jacobshade/lbuc-admin/server/model"

func CreateCheck(player model.Player, task model.Task) (model.Check, error) {
	// Create Check
	check := model.Check{PlayerID: player.ID, TaskID: task.ID, Checked: false}
	if err := DB.Create(&check).Error; err != nil {
		return model.Check{}, err
	}

	return check, nil
}

func GetCheck(playerID uint, taskID uint) (model.Check, error) {
	var check model.Check
	if err := DB.Where("player_id = ? AND task_id = ?", playerID, taskID).First(&check).Error; err != nil {
		return model.Check{}, err
	}
	return check, nil
}

func GetChecksForPlayer(player model.Player) ([]model.Check, error) {
	// Get Checks
	var checks []model.Check
	if err := DB.Where("player_id = ?", player.ID).Find(&checks).Error; err != nil {
		return []model.Check{}, err
	}

	return checks, nil
}

func GetChecksForTask(task model.Task) ([]model.Check, error) {
	var checks []model.Check
	if err := DB.Where("task_id = ?", task.ID).Find(&checks).Error; err != nil {
		return []model.Check{}, err
	}

	return checks, nil
}

func GetAllChecks() ([]model.Check, error) {
	var checks []model.Check
	if err := DB.Find(&checks).Error; err != nil {
		return []model.Check{}, err
	}
	return checks, nil
}

func UpdateCheck(check model.Check) error {
	if err := DB.Save(&check).Error; err != nil {
		return err
	}
	return nil
}

func DeleteCheck(check model.Check) error {
	if err := DB.Delete(&check).Error; err != nil {
		return err
	}
	return nil
}
