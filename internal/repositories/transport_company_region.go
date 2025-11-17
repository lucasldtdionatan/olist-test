package repositories

import (
	"context"
	"olist-project/internal/dto"
	"olist-project/internal/entities"

	"gorm.io/gorm"
)

type TransportCompanyRegionRepository interface {
	Create(region *entities.TransportCompanyRegion) error
	FindAll(filters dto.TransportCompanyRegionFilter) ([]entities.TransportCompanyRegion, error)
	FindByID(ctx context.Context, id uint) (*entities.TransportCompanyRegion, error)
	FindByName(ctx context.Context, name string) ([]entities.TransportCompanyRegionWithCompany, error)
	Delete(ctx context.Context, id uint) error
}

type transportCompanyRegionRepository struct {
	db *gorm.DB
}

func NewTransportCompanyRegionRepository(db *gorm.DB) TransportCompanyRegionRepository {
	return &transportCompanyRegionRepository{db}
}

func (r *transportCompanyRegionRepository) Create(region *entities.TransportCompanyRegion) error {
	return r.db.Create(region).Error
}

func (r *transportCompanyRegionRepository) FindAll(filters dto.TransportCompanyRegionFilter) ([]entities.TransportCompanyRegion, error) {
	var regions []entities.TransportCompanyRegion

	query := r.db.Model(&entities.TransportCompanyRegion{})

	if filters.TransportCompanyID != nil {
		query = query.Where("transport_company_id = ?", filters.TransportCompanyID)
	}

	if err := query.Find(&regions).Error; err != nil {
		return nil, err
	}

	return regions, nil
}

func (r *transportCompanyRegionRepository) FindByID(ctx context.Context, id uint) (*entities.TransportCompanyRegion, error) {
	var tcr entities.TransportCompanyRegion

	err := r.db.WithContext(ctx).First(&tcr, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &tcr, nil
}

func (r *transportCompanyRegionRepository) FindByName(ctx context.Context, name string) ([]entities.TransportCompanyRegionWithCompany, error) {
	var regions []entities.TransportCompanyRegionWithCompany

	err := r.db.WithContext(ctx).
		Model(&entities.TransportCompanyRegion{}).
		Preload("TransportCompany").
		Where("name = ?", name).
		Find(&regions).Error

	if err != nil {
		return nil, err
	}

	return regions, nil
}

func (r *transportCompanyRegionRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Delete(&entities.TransportCompanyRegion{}, "id = ?", id).Error
}
