package dto

import "github.com/google/uuid"

type ShippingQuoteRequest struct {
	DestinationState string  `json:"destination_state" binding:"required"`
	Weight           float64 `json:"weight" binding:"required,gt=0"`
}

type ShippingQuoteResponse struct {
	TransportCompanyID   uuid.UUID `json:"transport_company_id"`
	TransportCompanyName string    `json:"transport_company"`
	Price                float64   `json:"price"`
	EstimatedDays        int       `json:"estimated_days"`
}
