package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/model"
)

// Serialized Task
type Task struct {
	ID          uint   `json:"id"`
	Description string `json:"description"`
}

func CreateResponseTask(taskModel model.Task) Task {
	return Task{ID: taskModel.ID, Description: taskModel.Description}
}

func CreateTask(task Task, teamID uint) (model.Task, error) {
	newTask := model.Task{}
	newTask.Description = task.Description
	newTask.TeamRefer = teamID

	if err := database.DB.Create(&newTask).Error; err != nil {
		return model.Task{}, err
	}

	return newTask, nil
}

func GetTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	task := model.Task{}
	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	responseTask := CreateResponseTask(task)
	return c.Status(fiber.StatusOK).JSON(responseTask)
}

func UpdateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	task := model.Task{}
	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	var updateData Task
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	task.Description = updateData.Description
	database.DB.Save(&task)

	responseTask := CreateResponseTask(task)
	return c.Status(fiber.StatusOK).JSON(responseTask)
}

func DeleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	task := model.Task{}
	if err := database.DB.First(&task, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	database.DB.Delete(&task)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Task deleted"})
}
