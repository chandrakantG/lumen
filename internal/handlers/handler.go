package handlers

import (
	"github.com/chandrakantG/lumen/internal/services"
	"github.com/rs/zerolog"
)

type Handler struct {
	// Add any dependencies you need here, such as services or repositories
	logger  zerolog.Logger
	service services.Service
}

// NewHandler creates a new Handler instance
func NewHandler(logger zerolog.Logger, service services.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}
