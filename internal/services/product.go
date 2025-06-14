package services

import "github.com/chandrakantG/lumen/internal/models"

func (s service) GetRevenueByProduct(req models.ProductCreateRequest) ([]models.ProductRevenueResponse, error) {
	revenues, err := s.productRepo.GetRevenueByProduct(s.db, req.Start, req.End)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to get revenue by product")
		return nil, err
	}
	return revenues, nil
}
