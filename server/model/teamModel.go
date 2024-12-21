package model

type Team struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	TeamName string   `json:"team_name"`
	Players  []Player `gorm:"many2many:player_teams;"`
	Tasks    []Task   `gorm:"foreignKey:TeamRefer"`
}
