package repositories

import (
	"context"
	"olist-project/internal/entities"

	"gorm.io/gorm"
)

type ShipmentRepository interface {
	Create(ctx context.Context, s *entities.Shipment) error
	// FindByPackage(ctx context.Context, packageID uuid.UUID) (*entities.Shipment, error)
}

type shipmentRepository struct {
	db *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) ShipmentRepository {
	return &shipmentRepository{db}
}

func (r *shipmentRepository) Create(ctx context.Context, s *entities.Shipment) error {
	return r.db.Create(s).Error
}
