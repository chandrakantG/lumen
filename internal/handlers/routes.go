package handlers

import (
	"database/sql"
	"log"

	"github.com/chandrakantG/lumen/config"
	"github.com/chandrakantG/lumen/internal/database"
	"github.com/chandrakantG/lumen/internal/repositories"
	"github.com/chandrakantG/lumen/internal/services"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	// Initialize the database connection
	db := database.DB
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	handler := configureHandlers(db)

	r := gin.New()
	r.Use(gin.Recovery())
	v1Api := r.Group("/api/v1")

	v1Api.POST("/revenue/product", handler.GetRevenueByProduct)

	r.Run(":" + config.GetAppConfig().SERVER_PORT)
}

func configureHandlers(db *sql.DB) *Handler {

	logger := config.GetLogger()

	// Initialize repositories with the database connection
	productRepo := repositories.NewProductRepository()
	service := services.NewService(db,
		logger,
		productRepo)

	// Initialize handlers with the services
	return NewHandler(logger, service)
}
