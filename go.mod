module src

go 1.16

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/labstack/echo/v4 v4.7.0
	github.com/labstack/gommon v0.3.1
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	gopkg.in/ini.v1 v1.66.4
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
)

replace local.packages/handler => ./handler
