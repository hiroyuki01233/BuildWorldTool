package model

import (
	"time"
)

type UserProjectRelation struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	ProjectId uint      `gorm:"not null;"json:"project_id"`
	UserId    uint      `gorm:"not null;"json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
