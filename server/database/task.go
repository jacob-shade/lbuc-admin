package database

import (
	"fmt"

	"github.com/jacobshade/lbuc-admin/server/model"
)

func CreateTask(task model.Task) (model.Task, error) {
	if err := DB.Create(&task).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func GetTask(taskID uint) (model.Task, error) {
	var task model.Task
	if err := DB.First(&task, taskID).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	if err := DB.Find(&tasks).Error; err != nil {
		return []model.Task{}, err
	}
	return tasks, nil
}

func UpdateTask(task model.Task) error {
	if err := DB.Save(&task).Error; err != nil {
		return err
	}
	fmt.Println("Task updated:", task)
	return nil
}

func DeleteTask(task model.Task) error {
	if err := DB.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}
