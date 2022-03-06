package main

import (
	"log"
	"net/http"
	"src/controller"

	jwt "github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

type Claims struct {
	Username string `json:"username"`
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

func User(c echo.Context) error {
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// username := claims["username"].(string)
	// log.Println("test")
	// return c.String(http.StatusOK, "Welcome !"+username)
	return c.JSON(http.StatusOK, "okay!")
}

func myMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("before action")

		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println(http.StatusOK, "missing cookie")
				return nil
			}
			log.Println(http.StatusOK, "bad request1")
			return nil
		}

		tknStr := cookie.Value
		claims := Claims{}
		// return c.JSON(http.StatusOK, tknStr)
		tkn, err := jwt.ParseWithClaims(tknStr, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// log.Println(claims.Username)

		// if err != nil {
		// 	if err == jwt.ErrSignatureInvalid {
		// 		log.Println(http.StatusOK, "wrong token")
		// 		c.Error(err)
		// 		return nil
		// 	}
		// 	log.Println(http.StatusOK, err)
		// 	c.Error(err)
		// 	return nil
		// }
		if !tkn.Valid {
			log.Println(http.StatusOK, "wrong token2")
			return nil
		}

		// if err := next(c); err != nil {
		// 	c.Error(err)
		// 	return nil
		// }
		return nil
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(config))

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.POST("/login", controller.UserLogin)
	e.POST("/register", controller.UserRegister)

	e.GET("/test", controller.UserController)
	e.GET("/hello", hello())

	r := e.Group("/restricted")
	r.Use(middleware.CORSWithConfig(config))
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:token",
	}))
	r.GET("/user", User)

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
