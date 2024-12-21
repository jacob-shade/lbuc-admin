package model

type Check struct {
	PlayerID uint `gorm:"primaryKey" json:"playerID"`
	TaskID   uint `gorm:"primaryKey" json:"taskID"`
	Checked  bool `json:"checked"`
}
