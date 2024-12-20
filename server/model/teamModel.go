package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamName string   `json:"team_name"`
	Players  []Player `gorm:"many2many:player_teams;"`
	Tasks    []Task   `gorm:"many2many:team_tasks;"`
}
