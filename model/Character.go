package model

import (
	"time"
)

type Character struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	ProjectId uint      `gorm:"not null;" json:"project_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	SubTitle  string    `gorm:"type:varchar(255);" json:"sub_title"`
	Birthday  time.Time `json:"birthday"`
	MainText  string    `gorm:"type:mediumtext;" json:"main_text"`
	Idea      string    `gorm:"type:mediumtext;" json:"idea"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
