package main

import (
	"log"
	"strings"

	"github.com/chandrakantG/lumen/config"
	"github.com/chandrakantG/lumen/internal/database"
	"github.com/chandrakantG/lumen/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file for local development
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error while loading .env file: %v", err.Error())
	}

	// load config
	config.LoadAppConfig()
	config.LoadDBConfig()

	// Initialize database connection
	database.InitializeDB()
	defer database.CloseDB()

	// Initialize logger
	config.InitLog()

	if strings.ToLower(config.GetAppConfig().ENVIRONMENT) == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	// handlers
	handlers.InitializeRoutes()
}
