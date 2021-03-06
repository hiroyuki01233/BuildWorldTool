package controller

import (
	"encoding/json"
	"net/http"
	"src/model"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// most need project name and admin name
func GetProjectByProjectNameAndAdminName(c echo.Context) error {
	// jsonBody := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	// if err != nil {
	// 	log.Error("empty json body")
	// 	return c.JSON(http.StatusOK, "empty json body")
	// }

	project := model.Project{}
	result := project.GetByNameAndAdminName(c.QueryParam("projectName"), c.QueryParam("adminName"))
	return c.JSON(http.StatusOK, result)
}

func GetAllProjects(c echo.Context) error {
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(jwt.MapClaims)
	idInt, _ := strconv.Atoi(claims["id"].(string))
	var id uint = uint(idInt)

	result := model.GetAllProjectsById(id)

	return c.JSON(http.StatusOK, result)

}

// most need name as project name, and title
func CreateProject(c echo.Context) error {
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(jwt.MapClaims)
	idInt, _ := strconv.Atoi(claims["id"].(string))
	var id uint = uint(idInt)

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	var existsFlg bool
	project := model.Project{}
	existsFlg = project.IsExistsByProjectNameAndUserId(jsonBody["name"].(string), id)

	if existsFlg {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "the project name already exists",
		})
	}

	var ProjectData = model.Project{
		AdminId: id,
		Name:    jsonBody["name"].(string),
		Title:   jsonBody["title"].(string),
	}

	ProjectData.Create()

	return c.JSON(http.StatusOK, map[string]string{
		"projectName": jsonBody["name"].(string),
		"adminName":   claims["name"].(string),
	})
}

func UpdateProject(c echo.Context) error {
	pid := c.Param("id")
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	Project := model.Project{}
	Project.Update(pid, jsonBody["column"].(string), jsonBody["body"].(string))

	return c.JSON(http.StatusOK, Project)
}
