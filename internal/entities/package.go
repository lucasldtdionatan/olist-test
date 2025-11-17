package entities

import (
	"time"

	"github.com/google/uuid"
)

type Package struct {
	ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Status           string    `gorm:"not null;default:'pending'"`
	Product          string    `gorm:"not null"`
	Weight           float32   `gorm:"not null"`
	DestinationState string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}
