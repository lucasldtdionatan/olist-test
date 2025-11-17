package dto

import "olist-project/internal/entities"

type TransportCompanyRegionFilter struct {
	TransportCompanyID *string `form:"transport_company_id"`
}

type CreateTransportCompanyRegionRequest struct {
	Name          string  `json:"name" binding:"required"`
	EstimatedDays int     `json:"estimated_days" binding:"required"`
	PricePerKg    float64 `json:"price_per_kg" binding:"required"`
}

type RegionResponse struct {
	ID                 uint    `json:"id"`
	Name               string  `json:"name"`
	EstimatedDays      int     `json:"estimated_days"`
	PricePerKg         float64 `json:"price_per_kg"`
	TransportCompanyID string  `json:"transport_company_id"`
}

func ToTransportCompanyRegionResponse(p entities.TransportCompanyRegion) RegionResponse {
	return RegionResponse{
		ID:                 p.ID,
		Name:               p.Name,
		EstimatedDays:      p.EstimatedDays,
		PricePerKg:         p.PricePerKg,
		TransportCompanyID: p.TransportCompanyID.String(),
	}
}

func ToTransportCompanyRegionResponseList(list []entities.TransportCompanyRegion) []RegionResponse {
	result := make([]RegionResponse, len(list))
	for i, item := range list {
		result[i] = ToTransportCompanyRegionResponse(item)
	}
	return result
}
