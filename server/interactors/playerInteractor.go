package interactors

import (
	"errors"
	"fmt"

	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/model"
)

func CreatePlayer(player model.Player) (model.Player, error) {
	// Check if player attributes are valid
	if player.PlayerName == "" || player.NickName == "" || player.Grade == "" ||
		player.Birthday.IsZero() || player.PlayerEmail == "" ||
		player.ParentName == "" || player.ParentEmail == "" ||
		player.ParentNumber == "" || player.Relationship == "" ||
		player.Address == "" || player.MedicalNotes == "" {
		return model.Player{}, errors.New("invalid player attributes")
	}

	// Create player in database
	player, err := database.CreatePlayer(player)
	if err != nil {
		return model.Player{}, err
	}

	return player, nil
}

func GetPlayer(id uint) (model.Player, error) {
	// Check for zero/invalid ID
	if id == 0 {
		return model.Player{}, errors.New("invalid id: id cannot be zero")
	}

	// Get player from database
	player, err := database.GetPlayer(id)
	if err != nil {
		return model.Player{}, fmt.Errorf("failed to get player: %w", err)
	}

	// Verify player exists (not empty)
	if player.ID == 0 {
		return model.Player{}, errors.New("player not found")
	}

	return player, nil
}

func GetAllPlayers() ([]model.Player, error) {
	// Get all players from database
	players, err := database.GetAllPlayers()
	if err != nil {
		return []model.Player{}, fmt.Errorf("failed to get all players: %w", err)
	}

	return players, nil
}

func UpdatePlayer(player model.Player, id uint) error {
	// Check for zero/invalid ID
	if id == 0 {
		return errors.New("invalid id: id cannot be zero")
	}

	// Make sure player exists
	oldPlayer, err := GetPlayer(id)
	if err != nil {
		return fmt.Errorf("failed to get player: %w", err)
	}

	// Make sure player ID cannot be changed
	player.ID = oldPlayer.ID

	// Update player in database
	err = database.UpdatePlayer(player)
	if err != nil {
		return fmt.Errorf("failed to update player: %w", err)
	}

	return nil
}

func DeletePlayer(id uint) error {
	// Check for zero/invalid ID
	if id == 0 {
		return errors.New("invalid id: id cannot be zero")
	}

	// Make sure player exists
	player, err := GetPlayer(id)
	if err != nil {
		return fmt.Errorf("failed to get player: %w", err)
	}

	// Delete player from database
	err = database.DeletePlayer(player.ID)
	if err != nil {
		return fmt.Errorf("failed to delete player: %w", err)
	}

	return nil
}
