package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jacobshade/lbuc-admin/server/interactors"
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

func GetTask(c *fiber.Ctx) error {
	// Get task id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Get task
	task, err := interactors.GetTask(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	responseTask := CreateResponseTask(task)
	return c.Status(fiber.StatusOK).JSON(responseTask)
}

func UpdateTask(c *fiber.Ctx) error {
	// Get task id
	fmt.Println("UpdateTask called")
	id, err := c.ParamsInt("id")
	if err != nil {
		fmt.Println("Error getting task id:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Get update data
	var task Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid description"})
	}

	// Update task
	taskModel := model.Task{ID: uint(id), Description: task.Description}
	if err := interactors.UpdateTask(taskModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update task"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Task successfully updated"})
}

func DeleteTask(c *fiber.Ctx) error {
	// Get task id
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Delete task
	if err := interactors.DeleteTask(uint(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Task not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Task deleted"})
}
