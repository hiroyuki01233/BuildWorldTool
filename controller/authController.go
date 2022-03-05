package controller

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func LoginController(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// とりあえずのパスワード認証
	if username != "taro" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// トークン作成
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Taro"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
