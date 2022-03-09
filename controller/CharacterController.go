package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"src/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetAllCharactersByProjectId(c echo.Context) error {
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(jwt.MapClaims)
	idInt, _ := strconv.Atoi(claims["id"].(string))
	var id uint = uint(idInt)
	fmt.Println(id)

	pidInt, _ := strconv.Atoi(c.QueryParam("projectId"))
	var pid uint = uint(pidInt)
	result := model.GetAllCharactersByProjectId(pid)

	return c.JSON(http.StatusOK, result)

}

// most need name as project name, and title
func CreateCharacter(c echo.Context) error {
	// userInfo := c.Get("user").(*jwt.Token)
	// claims := userInfo.Claims.(jwt.MapClaims)
	// idInt, _ := strconv.Atoi(claims["id"].(string))
	// var id uint = uint(idInt)

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	pidInt, _ := strconv.Atoi(jsonBody["projectId"].(string))
	var pid uint = uint(pidInt)
	// return c.JSON(http.StatusOK, pid)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	var existsFlg bool
	character := model.Character{}
	// return c.JSON(http.StatusOK, pid)
	existsFlg = character.IsExistsByCharacterNameAndProjectId(jsonBody["name"].(string), pid)

	if existsFlg {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "the character name already exists",
		})
	}

	var CharacterData = model.Character{
		ProjectId: pid,
		Name:      jsonBody["name"].(string),
		FullName:  jsonBody["fullName"].(string),
		Birthday:  time.Now(),
	}

	CharacterData.Create()

	return c.JSON(http.StatusOK, CharacterData)
}
