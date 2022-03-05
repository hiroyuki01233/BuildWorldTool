package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"name" gorm:"type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p *User) FirstById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).First(&p)
}

func (p *User) FirstByUserName(username string) (tx *gorm.DB) {
	return DB.Where("user_name = ?", username).First(&p)
}

func (p *User) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

// all collums update
func (p *User) Save() (tx *gorm.DB) {
	return DB.Save(&p)
}

func (p *User) Updates() (tx *gorm.DB) {
	// db.Model(&product).Updates(Product{Name: "hoge", Price: 20})
	return DB.Model(&p).Updates(p)
}

func (p *User) Delete() (tx *gorm.DB) {
	return DB.Delete(&p)
}

func (p *User) DeleteById(id uint) (tx *gorm.DB) {
	return DB.Where("id = ?", id).Delete(&p)
}

func (p *User) IsExistsByUserName(username string) bool {
	var count int64
	DB.Where("user_name = ?", username).Find(&p).Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
	// if count > 0 {
	// 	log.Error("username already used")
	// 	resp := c.JSON(http.StatusConflict, helper.ErrorLog(http.StatusConflict, "username already used", "EXT_REF"))
	// 	return resp
	// }
}
