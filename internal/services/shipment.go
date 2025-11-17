package services

import (
	"context"
	"olist-project/internal"
	"olist-project/internal/dto"
	"olist-project/internal/entities"
	"olist-project/internal/repositories"
	"time"
)

type ShipmentService struct {
	repo repositories.ShipmentRepository
}

func NewShipmentService(r repositories.ShipmentRepository) *ShipmentService {
	return &ShipmentService{r}
}

func (s *ShipmentService) CreateShipment(ctx context.Context, req dto.CreateShipmentRequest) (*dto.ShipmentResponse, error) {

	estimated := time.Now().Add(time.Hour * 24 * time.Duration(req.EstimatedDays))

	shipment := entities.Shipment{
		PackageID:           req.PackageID,
		TransportCompanyID:  req.TransportCompanyID,
		Price:               req.Price,
		EstimatedDays:       req.EstimatedDays,
		EstimatedDeliveryAt: estimated,
		TrackingCode:        internal.GenerateTrackingCode(),
	}

	if err := s.repo.Create(ctx, &shipment); err != nil {
		return nil, err
	}

	return &dto.ShipmentResponse{
		ID:                  shipment.ID,
		PackageID:           shipment.PackageID,
		TransportCompanyID:  shipment.TransportCompanyID,
		Price:               shipment.Price,
		EstimatedDays:       shipment.EstimatedDays,
		EstimatedDeliveryAt: shipment.EstimatedDeliveryAt,
		TrackingCode:        shipment.TrackingCode,
	}, nil
}
