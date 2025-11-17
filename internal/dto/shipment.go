package dto

import (
	"time"

	"github.com/google/uuid"
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
