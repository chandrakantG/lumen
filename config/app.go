package config

import "os"

type App struct {
	SERVER_PORT string
	JWT_SECRET  string
	ENVIRONMENT string
}

var appConfig App

func LoadAppConfig() App {
	appConfig = App{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		ENVIRONMENT: os.Getenv("ENVIRONMENT"),
	}
	return appConfig
}

func GetAppConfig() App {
	return appConfig
}
