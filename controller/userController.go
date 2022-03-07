package controller

import (
	"net/http"
	"src/model"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(jwt.MapClaims)
	idInt, _ := strconv.Atoi(claims["id"].(string))
	var id uint = uint(idInt)
	user := model.User{}
	user.FirstById(id)

	response := map[string]string{
		"id": claims["id"].(string), "username": user.Name,
	}

	return c.JSON(http.StatusOK, response)
}
