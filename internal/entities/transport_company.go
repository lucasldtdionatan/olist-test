package entities

import (
	"time"

	"github.com/google/uuid"
)

type TransportCompany struct {
	ID        uuid.UUID                `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string                   `gorm:"not null"`
	Regions   []TransportCompanyRegion `gorm:"foreignKey:TransportCompanyID"`
	CreatedAt time.Time                `gorm:"autoCreateTime"`
	UpdatedAt time.Time                `gorm:"autoUpdateTime"`
}

type TransportCompanyRegion struct {
	ID                 uint      `gorm:"primaryKey;autoIncrement"`
	TransportCompanyID uuid.UUID `gorm:"type:uuid;not null"`
	Name               string    `gorm:"not null"`
	EstimatedDays      int       `gorm:"not null"`
	PricePerKg         float64   `gorm:"not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

type TransportCompanyRegionWithCompany struct {
	TransportCompanyRegion
	TransportCompany TransportCompany `gorm:"foreignKey:TransportCompanyID"`
}
