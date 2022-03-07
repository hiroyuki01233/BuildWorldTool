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

type Msg struct {
	Message string `json:"message"`
}

type Claims struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func UserLogin(c echo.Context) error {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		log.Error("empty json body")
		return nil
	}
	user := model.User{}
	user.FirstByName(jsonBody["name"].(string))

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(jsonBody["password"].(string)))
	if err != nil {
		return c.JSON(http.StatusOK, "password is wrong")
	}

	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["id"] = strconv.FormatUint(uint64(user.ID), 10)
	// claims["name"] = user.Name
	// claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// return c.JSON(http.StatusOK, user.Name)

	claims := &Claims{
		ID:   strconv.FormatUint(uint64(user.ID), 10),
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Secure = false
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func UserRegister(c echo.Context) error {
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
	var UserInfo = model.User{
		Name:     jsonBody["username"].(string),
		Password: string(hash),
	}
	if !existsFlg {
		UserInfo.Create()
	} else {
		msg := &Msg{Message: "the username already exists"}
		return c.JSON(http.StatusBadRequest, msg)
	}

	claims := &Claims{
		ID:   strconv.FormatUint(uint64(UserInfo.ID), 10),
		Name: UserInfo.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Secure = false
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
