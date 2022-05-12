package main

import (
	"net/http"
	"log"
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
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	var rawDb *gorm.DB
	rawDb, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	// Config.DB = rawDb
	if err != nil {
		log.Fatalf("Error connect db: ", err)
	}
	Config.DB = rawDb.Debug()

	defer Config.DB.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Wellcome!!!")
	})
    apiGroup := e.Group("/api")
    Routes.APIRouter(apiGroup)
	e.Logger.Fatal(e.Start(":1323"))
}