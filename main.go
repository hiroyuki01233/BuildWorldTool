package main

import (
	"net/http"
	"src/controller"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
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
		}),
	)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/login", controller.LoginController)
	e.POST("/register", controller.UserRegister)

	e.GET("/test", controller.UserController)
	e.GET("/hello", hello())

	r := e.Group("/restricted")
	// r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowCredentials: bool(true),
	// 	// AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	}
}

// package main

// import (
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"

// 	"local.packages/handler"
// )

// func main() {
// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.POST("/login", handler.Login())
// 	r := e.Group("/restricted")
// 	r.Use(middleware.JWT([]byte("secret")))
// 	r.GET("/welcome", handler.Restricted())

// 	e.Logger.Fatal(e.Start(":1323"))
// }
