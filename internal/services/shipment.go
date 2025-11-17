package services

import (
	"context"
	"olist-project/internal"
	"olist-project/internal/dto"
	"olist-project/internal/entities"
	"olist-project/internal/repositories"
	"time"

	"github.com/google/uuid"
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

func (s *ShipmentService) List(filters dto.ShipmentFilter) ([]entities.Shipment, error) {
	return s.repo.FindAll(filters)
}

func (s *ShipmentService) GetByID(ctx context.Context, id uuid.UUID) (*entities.Shipment, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ShipmentService) Update(ctx context.Context, id uuid.UUID, req dto.UpdateShipmentRequest) (*entities.Shipment, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *ShipmentService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
