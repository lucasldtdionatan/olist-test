package dto

import (
	"olist-project/internal/entities"
	"time"
)

type PackageFilter struct {
	Status  *string `form:"status"`
	Product *string `form:"product"`
}

type PackageResponse struct {
	ID               string    `json:"id"`
	Status           string    `json:"status"`
	Product          string    `json:"product"`
	Weight           float32   `json:"weight"`
	DestinationState string    `json:"destination_state"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type CreatePackageDTO struct {
	Product          string  `json:"product" validate:"required"`
	Weight           float32 `json:"weight" validate:"required"`
	DestinationState string  `json:"destination_state" validate:"required"`
}

type UpdatePackageRequest struct {
	Status           *string  `json:"status"`
	Product          *string  `json:"product"`
	Weight           *float32 `json:"peso_kg"`
	DestinationState *string  `json:"destination_state"`
}

func ToPackageResponse(p entities.Package) PackageResponse {
	return PackageResponse{
		ID:               p.ID.String(),
		Status:           p.Status,
		Product:          p.Product,
		Weight:           p.Weight,
		DestinationState: p.DestinationState,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}
}

func ToPackageResponseList(list []entities.Package) []PackageResponse {
	result := make([]PackageResponse, len(list))
	for i, item := range list {
		result[i] = ToPackageResponse(item)
	}
	return result
}
