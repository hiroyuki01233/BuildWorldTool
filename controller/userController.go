package controller

import (
	"net/http"
	"src/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

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
	product := model.User{}
	product.FirstById(id)

	return c.JSON(http.StatusOK, product)
	// return c.String(http.StatusOK, "Hello, World!")
}
