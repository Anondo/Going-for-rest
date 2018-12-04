package config

import (
	"github.com/spf13/viper"
)

type Database struct {
	Name     string
	Port     int
	Host     string
	Username string
	Password string
	DB       string
}

var db Database

func LoadDB() {
	db = Database{
		Name:     viper.GetString("database.name"),
		Port:     viper.GetInt("database.port"),
		Host:     viper.GetString("database.host"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DB:       viper.GetString("database.db"),
	}
}

func GetDB() *Database {
	return &db
}
