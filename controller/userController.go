package controller

import (
	"net/http"
	"src/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Index is index route for health
// func (hc *UserController) Index(c echo.Context) error {
// 	u := &User{
// 		Name:  "Jon",
// 		Email: "jon@labstack.com",
// 	}
// 	return c.JSON(http.StatusOK, u)
// }

func UserController(c echo.Context) error {
	i, _ := strconv.Atoi(c.QueryParam("id"))
	id := uint(i)
	user := model.User{}
	user.FirstById(id)

	return c.JSON(http.StatusOK, user)
	// return c.String(http.StatusOK, "Hello, World!")
}
