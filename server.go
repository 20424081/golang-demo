package main

import (
	// "net/http"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	"goapp/Config"
	"goapp/Routes"
	"goapp/Controllers"
	"github.com/labstack/echo/v4/middleware"
	"goapp/Utils"
)

var (
	err error
)

func main() {
	// ENV
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Database
	var rawDb *gorm.DB
	rawDb, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		log.Fatalf("Error connect db: ", err)
	}
	// Config.DB = rawDb
	Config.DB = rawDb.Debug()
	defer Config.DB.Close()

	e := echo.New()
	e.Validator = Config.Validator
	e.Static("/static", "Assets")
	e.Renderer = Config.Renderer

	var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &Utils.JwtCustomClaims{},
		SigningKey: []byte(viper.GetString("PRIVATE_KEY")),
	})

	e.GET("/", Controllers.HomeIndex)

    authGroup := e.Group("/api/auth")
    Routes.AuthRouter(authGroup)
	todoGroup := e.Group("/api/todos", IsLoggedIn)
    Routes.TodoRouter(todoGroup)

	e.Logger.Fatal(e.Start(":1323"))
}