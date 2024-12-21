package model

import (
	"time"
)

type Player struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
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
	Teams        []Team    `gorm:"many2many:player_teams;"`
}
