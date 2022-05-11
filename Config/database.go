package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}


func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     viper.GetString("DATABASE_HOST"),
		Port:     viper.GetInt("DATABASE_PORT"),
		DBName:   viper.GetString("DATABASE_NAME"),
		User:     viper.GetString("DATABASE_USERNAME"),
		Password: viper.GetString("DATABASE_PASSWORD"),
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}