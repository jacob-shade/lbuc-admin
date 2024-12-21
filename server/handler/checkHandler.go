package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/model"
)

// Serialized Check
type Check struct {
	PlayerID uint `json:"playerID"`
	TaskID   uint `json:"taskID"`
	Checked  bool `json:"checked"`
}

func CreateResponseCheck(checkModel model.Check) Check {
	return Check{PlayerID: checkModel.PlayerID, TaskID: checkModel.TaskID, Checked: checkModel.Checked}
}

func CreateCheck(check Check) (Check, error) {
	newCheck := model.Check{}
	newCheck.PlayerID = check.PlayerID
	newCheck.TaskID = check.TaskID
	newCheck.Checked = check.Checked

	if err := database.DB.Create(&newCheck).Error; err != nil {
		return Check{}, err
	}

	responseCheck := CreateResponseCheck(newCheck)
	return responseCheck, nil
}

func GetChecksForPlayer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var checks []model.Check
	if err := database.DB.Where("player_id = ?", id).Find(&checks).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Checks not found"})
	}

	responseChecks := []Check{}
	for _, check := range checks {
		responseChecks = append(responseChecks, CreateResponseCheck(check))
	}

	return c.Status(fiber.StatusOK).JSON(responseChecks)
}

func GetChecksForTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var checks []model.Check
	if err := database.DB.Where("task_id = ?", id).Find(&checks).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Checks not found"})
	}

	responseChecks := []Check{}
	for _, check := range checks {
		responseChecks = append(responseChecks, CreateResponseCheck(check))
	}

	return c.Status(fiber.StatusOK).JSON(responseChecks)
}

func UpdateCheck(c *fiber.Ctx) error {
	var check Check
	if err := c.BodyParser(&check); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	database.DB.Model(&model.Check{}).Where("task_id = ? AND player_id = ?", check.TaskID, check.PlayerID).Update("checked", check.Checked)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Check updated"})
}
