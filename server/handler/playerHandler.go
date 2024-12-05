package handler

import (
	"time"

	"github.com/jacobshade/lbuc-admin/server/database"
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
	player := model.Player{}
	if err := c.BodyParser(&player); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Create(&player)

	return c.Status(fiber.StatusOK).JSON(CreateResponsePlayer(player))
}

func GetAllPlayers(c *fiber.Ctx) error {
	players := []model.Player{}
	if err := database.DB.Find(&players).Error; err != nil {
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

	player := model.Player{}
	if err := database.DB.First(&player, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}
	responsePlayer := CreateResponsePlayer(player)

	return c.Status(fiber.StatusOK).JSON(responsePlayer)
}

func UpdatePlayer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	player := model.Player{}
	if err := database.DB.First(&player, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}

	type UpdatePlayer struct {
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

	var updateData UpdatePlayer
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	player.PlayerName = updateData.PlayerName
	player.NickName = updateData.NickName
	player.Pronouns = updateData.Pronouns
	player.Grade = updateData.Grade
	player.Birthday = updateData.Birthday
	player.PlayerEmail = updateData.PlayerEmail
	player.ParentName = updateData.ParentName
	player.ParentEmail = updateData.ParentEmail
	player.ParentNumber = updateData.ParentNumber
	player.Relationship = updateData.Relationship
	player.Address = updateData.Address
	player.MedicalNotes = updateData.MedicalNotes

	database.DB.Save(&player)
	responsePlayer := CreateResponsePlayer(player)

	return c.Status(fiber.StatusOK).JSON(responsePlayer)
}

func DeletePlayer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	player := model.Player{}
	if err := database.DB.First(&player, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Player not found"})
	}

	if err := database.DB.Delete(&player).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	responsePlayer := CreateResponsePlayer(player)

	return c.Status(fiber.StatusOK).JSON(responsePlayer)
}
