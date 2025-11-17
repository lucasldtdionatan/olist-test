package repositories

import (
	"context"
	"olist-project/internal/dto"
	"olist-project/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransportCompanyRepository interface {
	Create(pkg *entities.TransportCompany) error
	FindAll() ([]entities.TransportCompany, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.TransportCompany, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdateTransportCompanyRequest) (*entities.TransportCompany, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type transportCompanyRepository struct {
	db *gorm.DB
}

func NewTransportCompanyRepository(db *gorm.DB) TransportCompanyRepository {
	return &transportCompanyRepository{db}
}

func (r *transportCompanyRepository) Create(pkg *entities.TransportCompany) error {
	return r.db.Create(pkg).Error
}

func (r *transportCompanyRepository) FindAll() ([]entities.TransportCompany, error) {
	var transportCompanies []entities.TransportCompany

	if err := r.db.
		Preload("Regions").
		Find(&transportCompanies).Error; err != nil {
		return nil, err
	}

	return transportCompanies, nil
}

func (r *transportCompanyRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.TransportCompany, error) {
	var tc entities.TransportCompany

	err := r.db.WithContext(ctx).First(&tc, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &tc, nil
}

func (r *transportCompanyRepository) Update(ctx context.Context, id uuid.UUID, req dto.UpdateTransportCompanyRequest) (*entities.TransportCompany, error) {
	var tc entities.TransportCompany

	if err := r.db.WithContext(ctx).
		Model(&tc).
		Where("id = ?", id).
		Updates(req).
		Error; err != nil {
		return nil, err
	}

	r.db.WithContext(ctx).First(&tc, "id = ?", id)
	return &tc, nil
}

func (r *transportCompanyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Delete(&entities.TransportCompany{}, "id = ?", id).Error
}
