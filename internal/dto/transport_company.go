package dto

import (
	"olist-project/internal/entities"
	"time"
)

type TransportCompanyResponse struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Regions   []RegionResponse `json:"regions"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

type UpdateTransportCompanyRequest struct {
	Name string `json:"name"`
}

type CreateTransportCompanyRequest struct {
	Name string `json:"name" validate:"required"`
}

func ToTransportCompanyResponse(p entities.TransportCompany) TransportCompanyResponse {
	return TransportCompanyResponse{
		ID:        p.ID.String(),
		Name:      p.Name,
		Regions:   ToTransportCompanyRegionResponseList(p.Regions),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToTransportCompanyResponseList(list []entities.TransportCompany) []TransportCompanyResponse {
	result := make([]TransportCompanyResponse, len(list))
	for i, item := range list {
		result[i] = ToTransportCompanyResponse(item)
	}
	return result
}
