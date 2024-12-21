package model

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Description string `json:"description"`
	TeamRefer   uint   `json:"teamRefer"`
}
