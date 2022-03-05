package controller

import (
	"encoding/json"
	"net/http"
	"src/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c echo.Context) error {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Error("empty json body")
		return nil
	}
	user := model.User{}
	user.FirstByUserName(jsonBody["username"].(string))

	// return c.JSON(http.StatusOK, user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonBody["password"].(string)))
	if err != nil {
		return c.JSON(http.StatusOK, "password is wrong")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
