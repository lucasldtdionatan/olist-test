package entities

import (
	"time"

	"github.com/google/uuid"
)

type Shipment struct {
	ID                 uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PackageID          uuid.UUID `gorm:"type:uuid;not null"`
	TransportCompanyID uuid.UUID `gorm:"type:uuid;not null"`
	TrackingCode       string    `gorm:"size:30;uniqueIndex;not null"`

	Price               float64   `gorm:"not null"`
	EstimatedDays       int       `gorm:"not null"`
	EstimatedDeliveryAt time.Time `gorm:"not null"`

	CreatedAt time.Time
}
