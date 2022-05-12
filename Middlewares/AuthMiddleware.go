package Middlewares

import (
    "github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(viper.GetString("PRIVATE_KEY"),),
})