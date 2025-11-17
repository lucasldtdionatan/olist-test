package repositories

import (
	"context"
	"olist-project/internal/dto"
	"olist-project/internal/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShipmentRepository interface {
	Create(ctx context.Context, s *entities.Shipment) error
	FindAll(filters dto.ShipmentFilter) ([]entities.Shipment, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Shipment, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateShipmentRequest) (*entities.Shipment, error)
	Delete(ctx context.Context, id uuid.UUID) error
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

func (r *shipmentRepository) FindAll(filters dto.ShipmentFilter) ([]entities.Shipment, error) {
	var shipments []entities.Shipment

	query := r.db.Model(&entities.Shipment{})

	if filters.PackageID != nil {
		query = query.Where("package_id = ?", *filters.PackageID)
	}

	if filters.TransportCompanyID != nil {
		query = query.Where("transport_company_id = ?", *filters.TransportCompanyID)
	}

	if err := query.Find(&shipments).Error; err != nil {
		return nil, err
	}

	return shipments, nil
}

func (r *shipmentRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Shipment, error) {
	var shipment entities.Shipment

	err := r.db.WithContext(ctx).First(&shipment, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &shipment, nil
}

func (r *shipmentRepository) Update(ctx context.Context, id uuid.UUID, req dto.UpdateShipmentRequest) (*entities.Shipment, error) {
	var shipment entities.Shipment

	if err := r.db.WithContext(ctx).First(&shipment, "id = ?", id).Error; err != nil {
		return nil, err
	}

	updates := make(map[string]interface{})

	if req.Price != nil {
		updates["price"] = *req.Price
	}

	if req.EstimatedDays != nil {
		updates["estimated_days"] = *req.EstimatedDays
		updates["estimated_delivery_at"] = shipment.CreatedAt.Add(time.Hour * 24 * time.Duration(*req.EstimatedDays))
	}

	if len(updates) > 0 {
		if err := r.db.WithContext(ctx).
			Model(&shipment).
			Where("id = ?", id).
			Updates(updates).
			Error; err != nil {
			return nil, err
		}
	}

	r.db.WithContext(ctx).First(&shipment, "id = ?", id)
	return &shipment, nil
}

func (r *shipmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Delete(&entities.Shipment{}, "id = ?", id).Error
}
