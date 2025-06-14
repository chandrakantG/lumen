package services

import (
	"database/sql"

	"github.com/chandrakantG/lumen/internal/models"
	"github.com/chandrakantG/lumen/internal/repositories"
	"github.com/rs/zerolog"
)

type service struct {
	db          *sql.DB
	logger      zerolog.Logger
	productRepo repositories.ProductRepository
}
type Service interface {
	GetRevenueByProduct(req models.ProductCreateRequest) ([]models.ProductRevenueResponse, error)
}

func NewService(db *sql.DB,
	logger zerolog.Logger,
	productRepo repositories.ProductRepository) Service {
	return service{
		db:          db,
		logger:      logger,
		productRepo: productRepo,
	}
}
