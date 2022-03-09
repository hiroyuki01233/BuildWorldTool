package model

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	ProjectId uint      `gorm:"not null;" json:"project_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	FullName  string    `gorm:"type:varchar(255);not null;" json:"full_name"`
	SubTitle  string    `gorm:"type:varchar(255);" json:"sub_title"`
	Birthday  time.Time `json:"birthday"`
	MainText  string    `gorm:"type:mediumtext;" json:"main_text"`
	Idea      string    `gorm:"type:mediumtext;" json:"idea"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllCharactersByProjectId(pid uint) []Character {
	var characters []Character
	DB.Where("project_id = ?", pid).Find(&characters)
	return characters
}

func (c *Character) IsExistsByCharacterNameAndProjectId(name string, projectId uint) bool {
	var count int64
	DB.Where("name = ? and project_id = ?", name, projectId).Find(&c).Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

func (c *Character) Create() (tx *gorm.DB) {

	return DB.Create(&c)
}
