package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint      `gorm:"primary_key;AUTO_INCREMENT"`
	AdminId     uint      `gorm:"not null;" json:"admin_id"`
	Name        string    `gorm:"type:varchar(255);not null;" json:"name"`
	Title       string    `gorm:"type:varchar(255);not null;" json:"title"`
	SubTitle    string    `gorm:"type:varchar(255);" json:"sub_title"`
	Description string    `gorm:"type:text;" json:"description"`
	MainText    string    `gorm:"type:mediumtext;" json:"main_text"`
	Chronology  string    `gorm:"type:mediumtext;" json:"chronology"`
	Idea        string    `gorm:"type:mediumtext;" json:"idea"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectAndUser struct {
	ID          uint      `gorm:"primary_key;AUTO_INCREMENT"`
	AdminId     uint      `gorm:"not null;" json:"admin_id"`
	AdminName   string    `gorm:"not null;" json:"admin_name"`
	Name        string    `gorm:"type:varchar(255);not null;" json:"name"`
	Title       string    `gorm:"type:varchar(255);not null;" json:"title"`
	SubTitle    string    `gorm:"type:varchar(255);" json:"sub_title"`
	Description string    `gorm:"type:text;" json:"description"`
	MainText    string    `gorm:"type:mediumtext;" json:"main_text"`
	Chronology  string    `gorm:"type:mediumtext;" json:"chronology"`
	Idea        string    `gorm:"type:mediumtext;" json:"idea"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ResultUserAndProject struct {
	Id          uint
	UserName    string
	Name        string
	Title       string
	SubTitle    string
	Description string
	MainText    string
	Chronology  string
	Idea        string
}

func (p *Project) GetByNameAndAdminName(projectName string, adminName string) ResultUserAndProject {
	result := ResultUserAndProject{}
	DB.Model(&User{}).Select("users.name as user_name, projects.id, projects.name, projects.title, projects.sub_title, projects.description, projects.main_text, projects.chronology, projects.idea").Joins("left join projects on projects.admin_id = users.id").Where("projects.name = ? and users.name = ?", projectName, adminName).Scan(&result)
	return result
}

func (p *Project) Create() (tx *gorm.DB) {
	return DB.Create(&p)
}

func GetAllProjectsById(userId uint) []ProjectAndUser {
	var projects []ProjectAndUser
	DB.Model(&User{}).Select("users.name as admin_name, projects.name, projects.title, projects.sub_title, projects.description, projects.main_text, projects.chronology, projects.idea, projects.updated_at").Joins("left join projects on projects.admin_id = users.id").Where("projects.admin_id = ?", userId).Scan(&projects)
	// DB.Where("admin_id= ?", userId).Find(&projects)
	return projects
}

func (p *Project) IsExistsByProjectNameAndUserId(projectName string, userId uint) bool {
	var count int64
	DB.Where("name = ? and admin_id= ?", projectName, userId).Find(&p).Count(&count)
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

func (p *Project) Update(id string, clumName string, body string) (tx *gorm.DB) {
	// db.Model(&product).Updates(Product{Name: "hoge", Price: 20})
	return DB.Model(&p).Where("id = ?", id).Update(clumName, body)
}
