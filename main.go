package main

import (
	"net/http"
	"src/controller"

	jwt "github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

type Claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret")
var config = middleware.CORSConfig{
	AllowCredentials: true,
	AllowOrigins:     []string{"http://localhost:3000"},
	AllowHeaders: []string{
		echo.HeaderAccessControlAllowHeaders,
		echo.HeaderAccessControlAllowOrigin,
		echo.HeaderContentType,
		echo.HeaderContentLength,
		echo.HeaderAcceptEncoding,
		echo.HeaderXCSRFToken,
		echo.HeaderAuthorization,
		echo.HeaderOrigin,
		echo.HeaderAccept,
		echo.HeaderAccessControlAllowCredentials,
	},
	AllowMethods: []string{
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
		http.MethodPost,
		http.MethodDelete,
	},
	MaxAge: 86400,
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome !"+username)
}

func SetHeaderAllowOrigin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:3000")
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(config))
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(SetHeaderAllowOrigin)

	e.POST("/login", controller.UserLogin)
	e.POST("/register", controller.UserRegister)

	e.GET("/hello", hello())

	r := e.Group("/restricted")
	r.Use(middleware.CORSWithConfig(config))
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:token",
	}))

	r.GET("/user", controller.GetUser)
	r.POST("/project", controller.CreateProject)
	r.GET("/project", controller.GetProjectByProjectNameAndAdminName)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	}
}
