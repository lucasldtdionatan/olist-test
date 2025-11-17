package repositories

import (
	"context"
	"olist-project/internal/dto"
	"olist-project/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PackageRepository interface {
	Create(pkg *entities.Package) error
	FindAll(filters dto.PackageFilter) ([]entities.Package, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Package, error)
	Update(ctx context.Context, id uuid.UUID, req dto.UpdatePackageRequest) (*entities.Package, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type packageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) PackageRepository {
	return &packageRepository{db}
}

func (r *packageRepository) Create(pkg *entities.Package) error {
	return r.db.Create(pkg).Error
}

func (r *packageRepository) FindAll(filters dto.PackageFilter) ([]entities.Package, error) {
	var packages []entities.Package

	query := r.db.Model(&entities.Package{})

	if filters.Product != nil {
		query = query.Where("product ILIKE ?", "%"+*filters.Product+"%")
	}

	if filters.Status != nil {
		query = query.Where("status = ?", *filters.Status)
	}

	if err := query.Find(&packages).Error; err != nil {
		return nil, err
	}

	return packages, nil
}

func (r *packageRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Package, error) {
	var pkg entities.Package

	err := r.db.WithContext(ctx).First(&pkg, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &pkg, nil
}

func (r *packageRepository) Update(ctx context.Context, id uuid.UUID, req dto.UpdatePackageRequest) (*entities.Package, error) {
	var pkg entities.Package

	if err := r.db.WithContext(ctx).
		Model(&pkg).
		Where("id = ?", id).
		Updates(req).
		Error; err != nil {
		return nil, err
	}

	r.db.WithContext(ctx).First(&pkg, "id = ?", id)
	return &pkg, nil
}

func (r *packageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Delete(&entities.Package{}, "id = ?", id).Error
}
