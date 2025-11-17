package dto

import (
	"time"

	"github.com/google/uuid"
	"olist-project/internal/entities"
)

type CreateShipmentRequest struct {
	PackageID          uuid.UUID `json:"package_id" binding:"required"`
	TransportCompanyID uuid.UUID `json:"transport_company_id" binding:"required"`
	Price              float64   `json:"price" binding:"required,gt=0"`
	EstimatedDays      int       `json:"estimated_days" binding:"required,gt=0"`
}

type ShipmentResponse struct {
	ID                  uuid.UUID `json:"id"`
	PackageID           uuid.UUID `json:"package_id"`
	TransportCompanyID  uuid.UUID `json:"transport_company_id"`
	Price               float64   `json:"price"`
	EstimatedDays       int       `json:"estimated_days"`
	EstimatedDeliveryAt time.Time `json:"estimated_delivery_at"`
	TrackingCode        string    `json:"tracking_code"`
}

type ShipmentFilter struct {
	PackageID          *string `form:"package_id"`
	TransportCompanyID *string `form:"transport_company_id"`
}

type UpdateShipmentRequest struct {
	Price         *float64 `json:"price"`
	EstimatedDays *int     `json:"estimated_days"`
}

func ToShipmentResponse(s entities.Shipment) ShipmentResponse {
	return ShipmentResponse{
		ID:                  s.ID,
		PackageID:           s.PackageID,
		TransportCompanyID:  s.TransportCompanyID,
		Price:               s.Price,
		EstimatedDays:      s.EstimatedDays,
		EstimatedDeliveryAt: s.EstimatedDeliveryAt,
		TrackingCode:        s.TrackingCode,
	}
}

func ToShipmentResponseList(list []entities.Shipment) []ShipmentResponse {
	result := make([]ShipmentResponse, len(list))
	for i, item := range list {
		result[i] = ToShipmentResponse(item)
	}
	return result
}
