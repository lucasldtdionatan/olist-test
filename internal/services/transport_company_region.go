package services

import (
	"olist-project/internal/dto"
	"olist-project/internal/entities"
	"olist-project/internal/repositories"

	"github.com/google/uuid"
)

type TransportCompanyRegionService struct {
	repo repositories.TransportCompanyRegionRepository
}

func NewTransportCompanyRegionService(r repositories.TransportCompanyRegionRepository) *TransportCompanyRegionService {
	return &TransportCompanyRegionService{r}
}

func (s *TransportCompanyRegionService) CreateForCompany(companyID uuid.UUID, data dto.CreateTransportCompanyRegionRequest) (*entities.TransportCompanyRegion, error) {
	region := &entities.TransportCompanyRegion{
		Name:               data.Name,
		EstimatedDays:      data.EstimatedDays,
		PricePerKg:         data.PricePerKg,
		TransportCompanyID: companyID,
	}

	err := s.repo.Create(region)
	return region, err
}

func (s *TransportCompanyRegionService) ListForCompany(companyID uuid.UUID) ([]entities.TransportCompanyRegion, error) {
	idStr := companyID.String()

	filter := &dto.TransportCompanyRegionFilter{
		TransportCompanyID: &idStr,
	}

	return s.repo.FindAll(*filter)
}
