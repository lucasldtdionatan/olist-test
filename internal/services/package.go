package services

import (
	"context"
	"olist-project/internal/dto"
	"olist-project/internal/entities"
	"olist-project/internal/repositories"

	"github.com/google/uuid"
)

type PackageService struct {
	repo repositories.PackageRepository
}

func NewPackageService(r repositories.PackageRepository) *PackageService {
	return &PackageService{r}
}

func (s *PackageService) Create(data dto.CreatePackageDTO) (*entities.Package, error) {
	pkg := &entities.Package{
		Product:          data.Product,
		Weight:           data.Weight,
		DestinationState: data.DestinationState,
	}

	err := s.repo.Create(pkg)
	return pkg, err
}

func (s *PackageService) List(filters dto.PackageFilter) ([]entities.Package, error) {
	return s.repo.FindAll(filters)
}

func (s *PackageService) GetByID(ctx context.Context, id uuid.UUID) (*entities.Package, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *PackageService) Update(ctx context.Context, id uuid.UUID, req dto.UpdatePackageRequest) (*entities.Package, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *PackageService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
