package controller

import (
	"encoding/json"
	"net/http"
	"src/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// Index is index route for health
// func (hc *UserController) Index(c echo.Context) error {
// 	u := &User{
// 		Name:  "Jon",
// 		Email: "jon@labstack.com",
// 	}
// 	return c.JSON(http.StatusOK, u)
// }

type Msg struct {
	Message string `json:"message"`
}

func UserController(c echo.Context) error {
	i, _ := strconv.Atoi(c.QueryParam("id"))
	id := uint(i)
	user := model.User{}
	user.FirstById(id)

	return c.JSON(http.StatusOK, user)
	// return c.String(http.StatusOK, "Hello, World!")
}

func UserRegister(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:3000")
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Error("empty json body")
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(jsonBody["password"].(string)), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	var existsFlg bool
	user := model.User{}
	existsFlg = user.IsExistsByUserName(jsonBody["username"].(string))
	var insertUser = model.User{
		UserName: jsonBody["username"].(string),
		Password: string(hash),
	}
	if !existsFlg {
		insertUser.Create()
	} else {
		msg := &Msg{Message: "the username already exists"}
		return c.JSON(http.StatusBadRequest, msg)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.UserName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = t
	cookie.Secure = false
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"response": "ok",
	})
}
