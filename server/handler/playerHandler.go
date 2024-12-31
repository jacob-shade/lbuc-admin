package handler

import (
	"time"

	"github.com/jacobshade/lbuc-admin/server/interactors"
	"github.com/jacobshade/lbuc-admin/server/model"

	"github.com/gofiber/fiber/v2"
)

// Serialized Player
type Player struct {
	ID           uint      `json:"id"`
	PlayerName   string    `json:"player_name"`
	NickName     string    `json:"nick_name"`
	Pronouns     string    `json:"pronouns"`
	Grade        string    `json:"grade"`
	Birthday     time.Time `json:"birthday"`
	PlayerEmail  string    `json:"player_email"`
	ParentName   string    `json:"parent_name"`
	ParentEmail  string    `json:"parent_email"`
	ParentNumber string    `json:"parent_number"`
	Relationship string    `json:"relationship"`
	Address      string    `json:"address"`
	MedicalNotes string    `json:"medical_notes"`
}

func CreateResponsePlayer(playerModel model.Player) Player {
	return Player{ID: playerModel.ID, PlayerName: playerModel.PlayerName, NickName: playerModel.NickName,
		Pronouns: playerModel.Pronouns, Grade: playerModel.Grade,
		Birthday: playerModel.Birthday, PlayerEmail: playerModel.PlayerEmail,
		ParentName: playerModel.ParentName, ParentEmail: playerModel.ParentEmail,
		ParentNumber: playerModel.ParentNumber, Relationship: playerModel.Relationship,
		Address: playerModel.Address, MedicalNotes: playerModel.MedicalNotes}
}

func CreatePlayer(c *fiber.Ctx) error {
	var body model.Player
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdPlayer, err := interactors.CreatePlayer(body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(CreateResponsePlayer(createdPlayer))
}

func GetAllPlayers(c *fiber.Ctx) error {
	players, err := interactors.GetAllPlayers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	responsePlayers := []Player{}
	for _, player := range players {
		responsePlayers = append(responsePlayers, CreateResponsePlayer(player))
	}

	return c.Status(fiber.StatusOK).JSON(responsePlayers)
}

func GetPlayer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	player, err := interactors.GetPlayer(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	responsePlayer := CreateResponsePlayer(player)

	return c.Status(fiber.StatusOK).JSON(responsePlayer)
}

func UpdatePlayer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	player, err := interactors.GetPlayer(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var updateData model.Player
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = interactors.UpdatePlayer(updateData, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	responsePlayer := CreateResponsePlayer(player)

	return c.Status(fiber.StatusOK).JSON(responsePlayer)
}

func DeletePlayer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err = interactors.DeletePlayer(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Player deleted successfully"})
}
