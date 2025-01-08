package database

import "github.com/jacobshade/lbuc-admin/server/model"

func CreateTeam(team model.Team) (model.Team, error) {
	if err := DB.Create(&team).Error; err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func GetTeam(id uint) (model.Team, error) {
	var team model.Team
	if err := DB.Preload("Players").Preload("Tasks").First(&team, id).Error; err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func GetAllTeams() ([]model.Team, error) {
	var teams []model.Team
	if err := DB.Find(&teams).Error; err != nil {
		return []model.Team{}, err
	}
	return teams, nil
}

func GetAllPlayersOnTeam(teamID uint) ([]model.Player, error) {
	var players []model.Player
	if err := DB.Raw("SELECT * FROM player_teams WHERE team_id = ?", teamID).Scan(&players).Error; err != nil {
		return []model.Player{}, err
	}
	return players, nil
}

func UpdateTeam(team model.Team) error {
	if err := DB.Save(&team).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTeam(team model.Team) error {
	if err := DB.Delete(&team).Error; err != nil {
		return err
	}
	return nil
}

func AddPlayersToTeam(team model.Team, players []model.Player) error {
	if err := DB.Model(&team).Association("Players").Append(&players); err != nil {
		return err
	}
	return nil
}

func RemovePlayerFromTeam(team model.Team, player model.Player) error {
	if err := DB.Model(&team).Association("Players").Delete(&player); err != nil {
		return err
	}
	return nil
}

func RemoveTaskFromTeam(team model.Team, task model.Task) error {
	if err := DB.Model(&team).Association("Tasks").Delete(&task); err != nil {
		return err
	}
	return nil
}
