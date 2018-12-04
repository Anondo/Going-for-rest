package config

import "github.com/spf13/viper"

type App struct {
	Host string
	Port int
}

var app App

func LoadApp() {

	port := viper.GetInt("port")
	if port == 0 {
		port = viper.GetInt("app.port")
	}

	app = App{
		Host: viper.GetString("app.host"),
		Port: port,
	}
}

func GetApp() *App {
	return &app
}
