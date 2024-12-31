package database

import (
	"github.com/jacobshade/lbuc-admin/server/model"
)

func CreatePlayer(player model.Player) (model.Player, error) {
	if err := DB.Create(&player).Error; err != nil {
		return model.Player{}, err
	}
	return player, nil
}

func GetPlayer(id uint) (model.Player, error) {
	var player model.Player
	if err := DB.First(&player, id).Error; err != nil {
		return model.Player{}, err
	}
	return player, nil
}

func GetAllPlayers() ([]model.Player, error) {
	var players []model.Player
	if err := DB.Find(&players).Error; err != nil {
		return []model.Player{}, err
	}
	return players, nil
}

func UpdatePlayer(player model.Player) error {
	if err := DB.Save(&player).Error; err != nil {
		return err
	}
	return nil
}

func DeletePlayer(id uint) error {
	if err := DB.Delete(&model.Player{}, id).Error; err != nil {
		return err
	}
	return nil
}
