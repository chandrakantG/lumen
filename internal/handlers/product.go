package handlers

import (
	"github.com/chandrakantG/lumen/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRevenueByProduct(ctx *gin.Context) {
	h.logger.Info().Msg("get revenue by product api started")
	var req models.ProductCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error().Err(err).Msg("Failed to bind request")
		ctx.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	revenues, err := h.service.GetRevenueByProduct(req)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get revenue by product")
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(200, revenues)
	h.logger.Info().Msg("get revenue by product api completed")
}
