package services

import (
	"context"
	"olist-project/internal/dto"
	"olist-project/internal/entities"
	"olist-project/internal/repositories"

	"github.com/google/uuid"
)

type TransportCompanyService struct {
	repo repositories.TransportCompanyRepository
}

func NewTransportCompanyService(r repositories.TransportCompanyRepository) *TransportCompanyService {
	return &TransportCompanyService{r}
}

func (s *TransportCompanyService) Create(data dto.CreateTransportCompanyRequest) (*entities.TransportCompany, error) {
	tc := &entities.TransportCompany{
		Name: data.Name,
	}

	err := s.repo.Create(tc)
	return tc, err
}

func (s *TransportCompanyService) List() ([]entities.TransportCompany, error) {
	return s.repo.FindAll()
}

func (s *TransportCompanyService) GetByID(ctx context.Context, id uuid.UUID) (*entities.TransportCompany, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *TransportCompanyService) Update(ctx context.Context, id uuid.UUID, req dto.UpdateTransportCompanyRequest) (*entities.TransportCompany, error) {
	return s.repo.Update(ctx, id, req)
}

func (s *TransportCompanyService) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
