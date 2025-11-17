package services

import (
	"context"
	"olist-project/internal/dto"
	"olist-project/internal/repositories"
)

type ShippingService interface {
	GetQuotesByRegion(ctx context.Context, req dto.ShippingQuoteRequest) ([]dto.ShippingQuoteResponse, error)
}

type shippingService struct {
	regionRepo repositories.TransportCompanyRegionRepository
}

func NewShippingService(rRepo repositories.TransportCompanyRegionRepository) ShippingService {
	return &shippingService{regionRepo: rRepo}
}

func (s *shippingService) GetQuotesByRegion(
	ctx context.Context,
	req dto.ShippingQuoteRequest,
) ([]dto.ShippingQuoteResponse, error) {

	regions, err := s.regionRepo.FindByName(ctx, req.DestinationState)
	if err != nil {
		return nil, err
	}

	var results []dto.ShippingQuoteResponse

	for _, region := range regions {
		price := req.Weight * region.PricePerKg

		results = append(results, dto.ShippingQuoteResponse{
			TransportCompanyID:   region.TransportCompanyID,
			TransportCompanyName: region.TransportCompany.Name,
			EstimatedDays:        region.EstimatedDays,
			Price:                price,
		})
	}

	return results, nil
}
