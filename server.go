package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
	"goapp/Config"
	"goapp/Routes"
)

var (
	err error
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	var rawDb *gorm.DB
	rawDb, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	// Config.DB = rawDb
	Config.DB = rawDb.Debug()

	defer Config.DB.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
    apiGroup := e.Group("/api")
    Routes.APIRouter(apiGroup)
	e.Logger.Fatal(e.Start(":1323"))
}