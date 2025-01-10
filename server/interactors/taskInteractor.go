package interactors

import (
	"fmt"

	"github.com/jacobshade/lbuc-admin/server/database"
	"github.com/jacobshade/lbuc-admin/server/model"
)

func CreateTask(description string, team model.Team) (model.Task, error) {
	task := model.Task{Description: description, TeamRefer: team.ID}
	if _, err := database.CreateTask(task); err != nil {
		return model.Task{}, fmt.Errorf("failed to create task: %w", err)
	}
	return task, nil
}

func GetTask(taskID uint) (model.Task, error) {
	task, err := database.GetTask(taskID)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to get task: %w", err)
	}
	return task, nil
}

func GetAllTasks() ([]model.Task, error) {
	tasks, err := database.GetAllTasks()
	if err != nil {
		return []model.Task{}, fmt.Errorf("failed to get all tasks: %w", err)
	}
	return tasks, nil
}

func UpdateTask(task model.Task) error {
	// Get task from database
	oldTask, err := database.GetTask(task.ID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Update task in database
	task.ID = oldTask.ID
	if err := database.UpdateTask(task); err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

func DeleteTask(taskID uint) error {
	// Get task from database
	task, err := database.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Delete task from database
	if err := database.DeleteTask(task); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
